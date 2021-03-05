import requests
import json

if __name__=='__main__':
    with open('sample_objects/json/box_inside_walls.json') as json_file:
        response = requests.post('http://127.0.0.1:8000/path-tracing', json=json.load(json_file))
        print(response)
