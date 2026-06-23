import requests

BASE_URL = "http://localhost:8080"


# /products

def test_get_products_returns_200():
    response = requests.get(f"{BASE_URL}/products")
    assert response.status_code == 200


def test_get_products_returns_two_items():
    response = requests.get(f"{BASE_URL}/products")
    data = response.json()
    assert len(data) == 2
    assert data[0]["name"] == "Jablko"

# negatywny (niedozwolony OPTIONS)
def test_products_options_returns_no_data():
    response = requests.options(f"{BASE_URL}/products")
    assert response.status_code == 200
    assert response.text == ""

# /payments

def test_post_payment_returns_201():
    response = requests.post(f"{BASE_URL}/payments", json={"amount": 49.99})
    assert response.status_code == 201

# negatywny (niedozwolony GET)
def test_payments_get_not_allowed():
    response = requests.get(f"{BASE_URL}/payments")
    assert response.status_code != 201
    assert response.text == ""

# /cart

def test_post_cart_returns_201():
    cart = [{"id": 1, "name": "Jablko", "price": 1.2}]
    response = requests.post(f"{BASE_URL}/cart", json=cart)
    assert response.status_code == 201

# negatywny (niedozwolony GET)
def test_cart_get_not_allowed():
    response = requests.get(f"{BASE_URL}/cart")
    assert response.status_code != 201
    assert response.text == ""