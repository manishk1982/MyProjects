from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from webdriver_manager.chrome import ChromeDriverManager
from time import sleep

options = Options()
options.add_argument("--start-maximized")
driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
driver.get("https://admin-demo.nopcommerce.com/login")

driver.find_element("id", "Email").clear()
driver.find_element("id", "Email").send_keys("admin@yourstore.com")

driver.find_element("id", "Password").clear()
driver.find_element("id", "Password").send_keys("admin")

driver.find_element("xpath", "//button[@type='submit']").click()

expected_title = "Dashboard / nopCommerce administration"
actual_title = driver.title
if expected_title == actual_title:
    print(f"Test Passed -- found title as '{actual_title}'")
else:
    print(f"Test Failed -- Actual title: '{actual_title}', Expected title:'{expected_title}'")

driver.close()
