from typing import Dict, Optional, List
import pycountry
from geopy.geocoders import Nominatim
from geopy.exc import GeocoderTimedOut, GeocoderUnavailable
import time

class LocationService:
    def __init__(self):
        self.geolocator = Nominatim(user_agent="location_service")
        self.cache = {}
        self.location_relations = {
            'city_to_province': {},
            'city_to_country': {},
            'province_to_country': {},
            'country_to_provinces': {},
            'province_to_cities': {}
        }
        self.last_request_time = 0
        self.min_delay = 1.0  # seconds

    def _respect_rate_limit(self):
        """Handle rate limiting for API calls"""
        current_time = time.time()
        time_since_last_request = current_time - self.last_request_time
        if time_since_last_request < self.min_delay:
            time.sleep(self.min_delay - time_since_last_request)
        self.last_request_time = time.time()

    def _get_location_data(self, location: str) -> Optional[Dict]:
        """Get location data from Nominatim"""
        try:
            self._respect_rate_limit()
            geocode_result = self.geolocator.geocode(
                location,
                addressdetails=True,
                language='en'
            )
            
            if geocode_result and geocode_result.raw.get('address'):
                return geocode_result.raw['address']
            return None
        except (GeocoderTimedOut, GeocoderUnavailable) as e:
            print(f"Geocoding error for {location}: {str(e)}")
            return None

    def update_location_relations(self, city: str = None, province: str = None, country: str = None):
        """Update location relations in memory"""
        if city and province:
            self.location_relations['city_to_province'][city.lower()] = province
        if city and country:
            self.location_relations['city_to_country'][city.lower()] = country
        if province and country:
            self.location_relations['province_to_country'][province.lower()] = country
            
            # Update country to provinces mapping
            if country not in self.location_relations['country_to_provinces']:
                self.location_relations['country_to_provinces'][country] = set()
            self.location_relations['country_to_provinces'][country].add(province)
            
            # Update province to cities mapping
            if province not in self.location_relations['province_to_cities']:
                self.location_relations['province_to_cities'][province] = set()
            if city:
                self.location_relations['province_to_cities'][province].add(city)

    def get_country_by_city(self, city: str) -> Optional[str]:
        """Get country name from city"""
        if not city:
            return None
            
        city = city.lower().strip()
        
        # Check cache
        if city in self.location_relations['city_to_country']:
            return self.location_relations['city_to_country'][city]
            
        # Try geocoding
        location_data = self._get_location_data(city)
        if location_data:
            country = location_data.get('country')
            if country:
                self.update_location_relations(city=city, country=country)
                return country
                
        return None

    def get_province_by_city(self, city: str) -> Optional[str]:
        """Get province name from city"""
        if not city:
            return None
            
        city = city.lower().strip()
        
        # Check cache
        if city in self.location_relations['city_to_province']:
            return self.location_relations['city_to_province'][city]
            
        # Try geocoding
        location_data = self._get_location_data(city)
        if location_data:
            province = (
                location_data.get('state') or
                location_data.get('province') or
                location_data.get('region')
            )
            if province:
                self.update_location_relations(city=city, province=province)
                return province
                
        return None

    def get_country_by_province(self, province: str) -> Optional[str]:
        """Get country name from province"""
        if not province:
            return None
            
        province = province.lower().strip()
        
        # Check cache
        if province in self.location_relations['province_to_country']:
            return self.location_relations['province_to_country'][province]
            
        # Try geocoding
        location_data = self._get_location_data(province)
        if location_data:
            country = location_data.get('country')
            if country:
                self.update_location_relations(province=province, country=country)
                return country
                
        return None

    def get_cities_by_province(self, province: str) -> List[str]:
        """Get all cities in a province"""
        if not province:
            return []
            
        province = province.strip()
        
        # Check cache
        if province in self.location_relations['province_to_cities']:
            return list(self.location_relations['province_to_cities'][province])
            
        return []

    def get_provinces_by_country(self, country: str) -> List[str]:
        """Get all provinces in a country"""
        if not country:
            return []
            
        country = country.strip()
        
        # Check cache
        if country in self.location_relations['country_to_provinces']:
            return list(self.location_relations['country_to_provinces'][country])
            
        return []

    def check_location_type(self, location: str) -> Dict[str, bool]:
        """Check if a string is a country, province, or city"""
        if not location or not isinstance(location, str):
            return {
                'is_country': False,
                'is_province': False,
                'is_city': False,
                'raw_input': location
            }

        location = location.strip()
        
        # Check cache
        cache_key = location.lower()
        if cache_key in self.cache:
            return self.cache[cache_key]

        # Try geocoding
        location_data = self._get_location_data(location)
        if location_data:
            result = {
                'is_country': bool(location_data.get('country') and 
                                 location_data.get('country').lower() == location.lower()),
                'is_province': bool(location_data.get('state') or 
                                  location_data.get('province') or 
                                  location_data.get('region')),
                'is_city': bool(location_data.get('city') or 
                              location_data.get('town') or 
                              location_data.get('village')),
                'raw_input': location
            }
            
            # Update relations
            if result['is_city']:
                city = location
                province = (location_data.get('state') or 
                          location_data.get('province') or 
                          location_data.get('region'))
                country = location_data.get('country')
                self.update_location_relations(city, province, country)
                
            self.cache[cache_key] = result
            return result

        # Check country using pycountry as fallback
        try:
            is_country = bool(pycountry.countries.search_fuzzy(location))
            result = {
                'is_country': is_country,
                'is_province': False,
                'is_city': False,
                'raw_input': location
            }
            self.cache[cache_key] = result
            return result
        except LookupError:
            result = {
                'is_country': False,
                'is_province': False,
                'is_city': False,
                'raw_input': location
            }
            self.cache[cache_key] = result
            return result

# Example usage:
if __name__ == "__main__":
    location_service = LocationService()

    # Test locations
    test_locations = [
        "Jakarta",
        "Bogor",
        "Sleman",
        "Bandung",
        "Sukabumi"
    ]

    # Test all functions
    for location in test_locations:
        print(f"\nTesting location: {location}")
        
        # Check location type
        type_result = location_service.check_location_type(location)
        print(f"Location type check: {type_result}")
        
        # Get country by city
        if type_result['is_city']:
            country = location_service.get_country_by_city(location)
            province = location_service.get_province_by_city(location)
            print(f"City is in country: {country}")
            print(f"City is in province: {province}")
            
        # Get country by province
        if type_result['is_province']:
            country = location_service.get_country_by_province(location)
            cities = location_service.get_cities_by_province(location)
            print(f"Province is in country: {country}")
            print(f"Cities in province: {cities}")
            
        # Get provinces by country
        if type_result['is_country']:
            provinces = location_service.get_provinces_by_country(location)
            print(f"Provinces in country: {provinces}")