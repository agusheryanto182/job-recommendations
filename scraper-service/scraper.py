import requests
import re
from bs4 import BeautifulSoup
import time
from typing import List, Dict
from preprocessor import JobDataPreprocessor, save_to_json, save_to_csv

class LinkedInJobScraper:
    """Scraper untuk mengambil data lowongan kerja dari LinkedIn"""
    
    def __init__(self):
        self.base_url = "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search"
        self.detail_url = "https://www.linkedin.com/jobs-guest/jobs/api/jobPosting"
        
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
        }
        
        print("Initializing data preprocessor...")
        try:
            self.preprocessor = JobDataPreprocessor()
            print("Data preprocessor initialized successfully!")
        except Exception as e:
            print(f"Error initializing preprocessor: {str(e)}")
            raise
    
    def get_jobs(self, keywords: str, location: str, max_pages: int = 10) -> List[Dict]:
        """
        Mengambil data pekerjaan dari LinkedIn berdasarkan keywords dan lokasi
        
        Args:
            keywords: Kata kunci pencarian
            location: Lokasi pekerjaan
            max_pages: Maksimum halaman yang akan di-scrape (default: 10)
            
        Returns:
            List of processed job dictionaries
        """
        all_jobs = []
        
        print(f"\nStarting job search for: {keywords} in {location}")
        print("This might take a while...\n")
        
        for page in range(max_pages):
            start = page * 25  # LinkedIn menampilkan 25 job per halaman
            
            params = {
                'keywords': keywords,
                'location': location,
                'start': start
            }
            
            try:
                print(f"Fetching page {page + 1}...")
                response = requests.get(
                    self.base_url,
                    headers=self.headers,
                    params=params,
                    timeout=30
                )
                
                if response.status_code == 200:
                    soup = BeautifulSoup(response.text, 'html.parser')
                    jobs = self._extract_jobs(soup)
                    
                    if not jobs:  # Tidak ada lagi pekerjaan yang ditemukan
                        print(f"No more jobs found on page {page + 1}")
                        break
                        
                    all_jobs.extend(jobs)
                    print(f"Successfully scraped page {page + 1}, Total jobs: {len(all_jobs)}")
                    
                    # Delay untuk menghindari rate limiting
                    time.sleep(1)
                else:
                    print(f"Failed to fetch page {page + 1}: Status code {response.status_code}")
                    break
                    
            except requests.exceptions.RequestException as e:
                print(f"Network error on page {page + 1}: {str(e)}")
                break
            except Exception as e:
                print(f"Unexpected error on page {page + 1}: {str(e)}")
                break
                
        return all_jobs

    def _get_list_ids(self, soup: BeautifulSoup) -> List[str]:
        id_list = []
        page_jobs = soup.find_all("li")
        for job in page_jobs:
            base_card_div = job.find("div", {"class": "base-card"})
            job_id = base_card_div.get("data-entity-urn").split(":")[3]
            id_list.append(job_id)
        return id_list
    
    def _extract_jobs(self, soup: BeautifulSoup) -> List[Dict]:
        jobs = []
        
        id_list = self._get_list_ids(soup)
        for job_id in id_list:
            try:
                job_url = f"https://www.linkedin.com/jobs-guest/jobs/api/jobPosting/{job_id}"
                job_response = requests.get(job_url)
                job_soup = BeautifulSoup(job_response.text, 'html.parser')
                
                # Extract basic job information
                raw_job = {
                    'location': job_soup.find('span', {'class': 'topcard__flavor topcard__flavor--bullet'}).text.strip(),
                    'posted_date':  job_soup.find("span", {"class": "posted-time-ago__text topcard__flavor--metadata"}).text.strip(),
                    'description': job_soup.find("div", {"show-more-less-html__markup show-more-less-html__markup--clamp-after-5 relative overflow-hidden"}).text.strip(),
                }
                
                criteria_list = job_soup.find("ul", {"class": "description__job-criteria-list"})
                if criteria_list:
                    criteria_items = criteria_list.find_all("li", {"class": "description__job-criteria-item"})
                    if len(criteria_items) >= 4:
                        spans = [item.find("span", {"class": ["description__job-criteria-text", "description__job-criteria-text--criteria"]}) 
                                for item in criteria_items]
                        
                        if all(spans):
                            raw_job['seniority_level'] = spans[0].text.strip()
                            raw_job['employment_level'] = spans[1].text.strip()
                            raw_job['job_function'] = spans[2].text.strip()
                            raw_job['industries'] = spans[3].text.strip()
                else:
                    raw_job['seniority_level'] = "Entry level"
                    raw_job['employment_level'] = "Full-time"
                    raw_job['job_function'] = "Engineering and Information Technology"
                    raw_job['industries'] = "Information Technology & Services"
                    
                # Process job data using our preprocessor
                processed_job = self.preprocessor.preprocess_job(raw_job)
                
                # Structure the final job data
                final_job = {
                    "job_id": job_id,
                    'title': job_soup.find("h2", {"class":"top-card-layout__title font-sans text-lg papabear:text-xl font-bold leading-open text-color-text mb-0 topcard__title"}).text.strip(),
                    'company': job_soup.find("a", {"class": "topcard__org-name-link topcard__flavor--black-link"}).text.strip(),
                    'location': processed_job['location'],
                    'seniority_level': processed_job['seniority_level'],
                    'employment_level': processed_job['employment_level'],
                    'job_function': processed_job['job_function'],
                    'industries': processed_job['industries'],
                    'link': job_soup.find('a', {'class': 'topcard__link'}).get('href'),
                    'posted_date': processed_job['posted_date'],
                    'processed_text': processed_job['processed_text']
                }
                
                try:
                    final_job['num_applicants'] = job_soup.find("span", {"class": "num-applicants__caption topcard__flavor--metadata topcard__flavor--bullet"}).text.strip()
                except:
                    final_job['num_applicants'] = job_soup.find("figcaption", {"class": "num-applicants__caption"}).text.strip()
                
                print(f"Processing job: {final_job['title']}")
                
                jobs.append(final_job)
                
                # Be nice to LinkedIn's servers
                time.sleep(2)
                
            except AttributeError as e:
                print(f"Error extracting job data: {str(e)}")
                continue
            except Exception as e:
                print(f"Unexpected error processing job: {str(e)}")
                continue
                
        return jobs

def main():
    """Main function untuk menjalankan scraper"""
    print("LinkedIn Job Scraper Starting...")
    print("Setting up dependencies...")
    
    # Verify dependencies
    try:
        import nltk
        import spacy
        import requests
        from bs4 import BeautifulSoup
    except ImportError as e:
        print(f"Error: Missing dependency - {str(e)}")
        print("Please install required packages:")
        print("pip install nltk spacy requests beautifulsoup4")
        print("python -m spacy download en_core_web_sm")
        return
    
    try:
        # Initialize scraper
        scraper = LinkedInJobScraper()
        
        # Get user input
        keywords = input("Enter job keywords (default: 'developer'): ") or "developer"
        location = input("Enter location (default: 'indonesia'): ") or "indonesia"
        max_pages = int(input("Enter maximum pages to scrape (default: 1): ") or "1")
        
        # Start scraping
        jobs = scraper.get_jobs(
            keywords=keywords,
            location=location,
            max_pages=max_pages
        )
        
        if jobs:
            # Save results in both formats
            timestamp = time.strftime("%Y%m%d_%H%M%S")
            json_filename = f'linkedin_jobs_{timestamp}.json'
            csv_filename = f'linkedin_jobs_{timestamp}.csv'
            
            save_to_json(jobs, json_filename)
            save_to_csv(jobs, csv_filename)
            
            print(f"\nSuccessfully saved {len(jobs)} jobs to:")
            print(f"- {json_filename} (complete data)")
            print(f"- {csv_filename} (flattened data for analysis)")
        else:
            print("\nNo jobs were found to save")
            
        print(f"\nTotal jobs scraped: {len(jobs)}")
        
    except Exception as e:
        print(f"\nError running scraper: {str(e)}")
        print("If this is a dependency error, please ensure all required packages are installed:")
        print("pip install nltk spacy requests beautifulsoup4")
        print("python -m spacy download en_core_web_sm")

if __name__ == "__main__":
    main()