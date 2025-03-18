from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import pandas as pd
import time

class LinkedInScraper:
    def __init__(self):
        self.setup_driver()
        
    def setup_driver(self):
        """Setup Chromium WebDriver"""
        options = Options()
        options.binary_location = "/usr/bin/chromium"
        options.add_argument('--headless')
        options.add_argument('--no-sandbox')
        options.add_argument('--disable-dev-shm-usage')
        options.add_argument('--window-size=1920,1080')
        
        service = Service('/usr/bin/chromedriver')
        self.driver = webdriver.Chrome(service=service, options=options)
        print("Browser setup successful!")

    def scrape_jobs(self, keywords, location, num_jobs=25):
        """Scrape job listings from LinkedIn"""
        try:
            # Format keywords and location for URL
            keywords = keywords.replace(" ", "%20")
            location = location.replace(" ", "%20")
            url = f"https://www.linkedin.com/jobs/search/?keywords={keywords}&location={location}"
            print(f"Accessing: {url}")
            
            self.driver.get(url)
            time.sleep(5)  # Wait longer for page to load
            
            # Wait for job list to be present
            print("Waiting for jobs to load...")
            WebDriverWait(self.driver, 10).until(
                EC.presence_of_element_located((By.CLASS_NAME, "jobs-search__results-list"))
            )
            
            # Get initial job count
            jobs_list = self.driver.find_element(By.CLASS_NAME, "jobs-search__results-list")
            jobs = jobs_list.find_elements(By.TAG_NAME, "li")
            print(f"Found {len(jobs)} jobs initially")
            
            # Scroll to load more jobs
            jobs = self._scroll_and_load_jobs(num_jobs)
            
            if jobs:
                return self._extract_job_details(jobs)
            return pd.DataFrame()
            
        except Exception as e:
            print(f"Error during scraping: {str(e)}")
            return pd.DataFrame()

    def _scroll_and_load_jobs(self, num_jobs):
        """Scroll page to load more jobs"""
        print(f"Attempting to load {num_jobs} jobs...")
        current_jobs = 0
        max_scrolls = (num_jobs // 25) + 1
        
        for scroll in range(max_scrolls):
            print(f"Scroll {scroll + 1}/{max_scrolls}")
            
            # Scroll to bottom
            self.driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
            time.sleep(2)
            
            try:
                # Try to click "Show more jobs" button
                button = self.driver.find_element(By.CSS_SELECTOR, "button.infinite-scroller__show-more-button")
                self.driver.execute_script("arguments[0].click();", button)
                print("Clicked 'Show more jobs' button")
                time.sleep(2)
            except:
                print("No 'Show more' button found")
            
            # Get current job count
            jobs_list = self.driver.find_element(By.CLASS_NAME, "jobs-search__results-list")
            jobs = jobs_list.find_elements(By.TAG_NAME, "li")
            current_jobs = len(jobs)
            print(f"Currently loaded: {current_jobs} jobs")
            
            if current_jobs >= num_jobs:
                break
                
        return jobs

    def _extract_job_details(self, job_cards):
        """Extract information from job cards"""
        print("Extracting job details...")
        jobs_data = []
        
        for i, card in enumerate(job_cards, 1):
            try:
                # Extract basic information
                title = card.find_element(By.CSS_SELECTOR, "h3").text
                company = card.find_element(By.CSS_SELECTOR, "h4").text
                location = card.find_element(By.CLASS_NAME, "job-search-card__location").text
                link = card.find_element(By.CSS_SELECTOR, "a").get_attribute('href')
                
                # Get posting date if available
                try:
                    date = card.find_element(By.CSS_SELECTOR, "time").get_attribute('datetime')
                except:
                    date = "Not specified"
                
                jobs_data.append({
                    'Title': title,
                    'Company': company,
                    'Location': location,
                    'Posted Date': date,
                    'Link': link
                })
                
                print(f"Processed job {i}: {title} at {company}")
                
            except Exception as e:
                print(f"Error extracting job {i}: {str(e)}")
                continue
                
        return pd.DataFrame(jobs_data)

    def save_data(self, df, filename):
        """Save data to CSV"""
        df.to_csv(f"data/{filename}", index=False)
        print(f"Data saved to data/{filename}")

    def close(self):
        """Close browser"""
        self.driver.quit()
        print("Browser closed")
        
    def test_connection(self):
        """Test connection to LinkedIn"""
        try:
            self.driver.get("https://www.linkedin.com")
            print("Connection successful!")
            return True
        except Exception as e:
            print(f"Connection failed: {str(e)}")
            return False