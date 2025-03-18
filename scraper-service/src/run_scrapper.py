from scrapper import LinkedInScraper
import time

def main():
    # Initialize scraper
    scraper = LinkedInScraper()
    
    try:
        # Scrape jobs
        print("\nStarting job search...")
        df = scraper.scrape_jobs(
            keywords="data scientist",
            location="Indonesia",
            num_jobs=25  # Start with small number for testing
        )
        
        # Save and display results
        if not df.empty:
            scraper.save_data(df, "linkedin_jobs.csv")
            print("\nResults Summary:")
            print("-" * 50)
            print(f"Total jobs found: {len(df)}")
            print("\nUnique Companies:")
            print(df['Company'].value_counts().head())
            print("\nLocations:")
            print(df['Location'].value_counts().head())
            print("\nSample Listings:")
            print(df[['Title', 'Company', 'Location']].head())
        else:
            print("\nNo jobs found! Please check the search criteria.")
            
    except Exception as e:
        print(f"\nError occurred: {str(e)}")
        
    finally:
        scraper.close()

if __name__ == "__main__":
    main()