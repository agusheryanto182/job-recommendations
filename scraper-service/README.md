# LinkedIn Job Scraper

Tool untuk mengumpulkan data lowongan kerja dari LinkedIn.

## Setup

1. Install dependencies:

   ```bash
   pip install -r requirements.txt
   ```

2. Jalankan scraper:
   ```python
   from src.scraper import LinkedInScraper
   scraper = LinkedInScraper()
   df = scraper.scrape_jobs("data scientist", "Indonesia")
   ```

## Features

- Scraping lowongan kerja LinkedIn
- Cleaning dan processing data
- Analisis data dasar
