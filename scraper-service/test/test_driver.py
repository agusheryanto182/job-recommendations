from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from webdriver_manager.chrome import ChromeDriverManager

def test_chrome_driver():
    try:
        print("Setting up Chrome options...")
        options = Options()
        options.add_argument('--no-sandbox')
        options.add_argument('--disable-dev-shm-usage')
        
        print("Installing ChromeDriver...")
        service = Service(ChromeDriverManager().install())
        
        print("Initializing Chrome Driver...")
        driver = webdriver.Chrome(service=service, options=options)
        
        print("Opening test page...")
        driver.get("https://www.google.com")
        
        print("Test successful!")
        driver.quit()
        
    except Exception as e:
        print(f"Test failed with error: {str(e)}")

if __name__ == "__main__":
    test_chrome_driver()