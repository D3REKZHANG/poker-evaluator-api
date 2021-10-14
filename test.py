import requests

print(requests.post('http://localhost:1323/rankHand', json={'cards': ['2h', '4s', '5d', 'Kh', 'As']}).text)

river = ['2h', '4s', '5d', 'Kh', 'As']
holes = [
    ['2s', '2d'],
    ['3s', '6d'],
]
#print(requests.post('http://localhost:1323/rankTable', json={'river': river, 'holes': holes }).text)
