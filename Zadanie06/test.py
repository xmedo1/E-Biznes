import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
import time

class Test:

    @pytest.fixture(scope="class")
    def driver(self):
        options = webdriver.ChromeOptions()
        options.add_argument("--start-maximized")
        options.add_argument("--incognito")
        options.add_experimental_option("excludeSwitches", ["enable-automation"])
        driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
        yield driver
        driver.quit()

    def test_01_page_load(self, driver):
        driver.get("https://www.saucedemo.com/")
        assert "Swag Labs" in driver.title

    def test_02_username_field_visible(self, driver):
        assert driver.find_element(By.ID, "user-name").is_displayed()

    def test_03_password_field_visible(self, driver):
        assert driver.find_element(By.ID, "password").is_displayed()

    def test_04_invalid_login(self, driver):
        driver.find_element(By.ID, "user-name").send_keys("error_user")
        driver.find_element(By.ID, "password").send_keys("wrong_password")
        driver.find_element(By.ID, "login-button").click()
        assert driver.find_element(By.CSS_SELECTOR, "[data-test='error']").is_displayed()

    def test_05_valid_login(self, driver):
        driver.refresh()
        driver.find_element(By.ID, "user-name").send_keys("standard_user")
        driver.find_element(By.ID, "password").send_keys("secret_sauce")
        driver.find_element(By.ID, "login-button").click()
        assert "inventory.html" in driver.current_url

    def test_06_burger_menu_opens(self, driver):
        driver.find_element(By.ID, "react-burger-menu-btn").click()
        time.sleep(0.5)
        assert driver.find_element(By.ID, "logout_sidebar_link").is_displayed()

    def test_07_burger_menu_closes(self, driver):
        driver.find_element(By.ID, "react-burger-cross-btn").click()
        time.sleep(0.5)
        assert "inventory.html" in driver.current_url

    def test_08_inventory_items_present(self, driver):
        items = driver.find_elements(By.CLASS_NAME, "inventory_item")
        assert len(items) == 6

    def test_09_sort_dropdown_visible(self, driver):
        assert driver.find_element(By.CLASS_NAME, "product_sort_container").is_displayed()

    def test_10_add_backpack_to_cart(self, driver):
        driver.find_element(By.ID, "add-to-cart-sauce-labs-backpack").click()
        assert driver.find_element(By.CLASS_NAME, "shopping_cart_badge").text == "1"

    def test_11_remove_backpack_from_cart(self, driver):
        driver.find_element(By.ID, "remove-sauce-labs-backpack").click()
        assert len(driver.find_elements(By.CLASS_NAME, "shopping_cart_badge")) == 0

    def test_12_add_bike_light_to_cart(self, driver):
        driver.find_element(By.ID, "add-to-cart-sauce-labs-bike-light").click()
        assert driver.find_element(By.CLASS_NAME, "shopping_cart_badge").text == "1"

    def test_13_open_cart_page(self, driver):
        driver.find_element(By.CLASS_NAME, "shopping_cart_link").click()
        assert "cart.html" in driver.current_url

    def test_14_continue_shopping(self, driver):
        driver.find_element(By.ID, "continue-shopping").click()
        assert "inventory.html" in driver.current_url

    def test_15_go_to_product_details(self, driver):
        driver.find_element(By.ID, "item_4_title_link").click()
        assert "inventory-item.html" in driver.current_url

    def test_16_back_to_products(self, driver):
        driver.find_element(By.ID, "back-to-products").click()
        assert "inventory.html" in driver.current_url

    def test_17_go_to_checkout(self, driver):
        driver.find_element(By.CLASS_NAME, "shopping_cart_link").click()
        driver.find_element(By.ID, "checkout").click()
        assert "checkout-step-one.html" in driver.current_url

    def test_18_cancel_checkout(self, driver):
        driver.find_element(By.ID, "cancel").click()
        assert "cart.html" in driver.current_url

    def test_19_footer_visible(self, driver):
        assert driver.find_element(By.CLASS_NAME, "footer").is_displayed()

    def test_20_logout(self, driver):
        driver.execute_script("window.scrollTo(0, 0);")
        driver.find_element(By.ID, "react-burger-menu-btn").click()
        time.sleep(1)
        logout_link = driver.find_element(By.ID, "logout_sidebar_link")
        driver.execute_script("arguments[0].click();", logout_link)
        assert driver.current_url == "https://www.saucedemo.com/"