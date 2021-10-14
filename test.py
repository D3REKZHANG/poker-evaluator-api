import requests

print(requests.post('http://localhost:1323/rankHand', json={'cards': ['2h', '4s', '5d', 'Kh', 'As']}).text)
