from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options

def test_selenium():
    try:
        print("Setting up Chromium options...")
        options = Options()
        options.binary_location = "/usr/bin/chromium"  # Specify Chromium binary
        options.add_argument('--headless')
        options.add_argument('--no-sandbox')
        options.add_argument('--disable-dev-shm-usage')
        
        print("Setting up ChromeDriver...")
        service = Service('/usr/bin/chromedriver')
        
        print("Starting browser...")
        driver = webdriver.Chrome(service=service, options=options)
        
        print("Testing connection...")
        driver.get("https://www.google.com")
        print("Test successful!")
        
        driver.quit()
        print("Browser closed")
        
    except Exception as e:
        print(f"Error: {str(e)}")

if __name__ == "__main__":
    test_selenium()