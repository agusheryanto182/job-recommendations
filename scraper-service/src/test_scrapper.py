from scrapper import LinkedInScraper

def test_setup():
    try:
        print("Initializing scraper...")
        scraper = LinkedInScraper()
        
        print("\nTesting connection...")
        if scraper.test_connection():
            print("All tests passed!")
        
        scraper.driver.quit()
        print("Browser closed")
        
    except Exception as e:
        print(f"Test failed: {e}")

if __name__ == "__main__":
    test_setup()