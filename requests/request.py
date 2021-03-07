import json

import uuid
from PIL import Image
import numpy as np
import requests

if __name__ == '__main__':
    response = dict()
    with open('../sample_objects/json/box_inside_walls.json') as json_file:
        response = requests.post('http://127.0.0.1:8081/path-tracing', json=json.load(json_file))
    try:
        colorMatrix = response.json()["Colors"]
        width = len(colorMatrix[0])
        height = len(colorMatrix)
        imageData = np.zeros((height, width, 3), dtype=np.uint8)
        for lineIndex in range(height):
            for columnIndex in range(width):
                for colorIndex in range(3):
                    imageData[lineIndex][columnIndex][colorIndex] = colorMatrix[lineIndex][columnIndex][colorIndex]
        image = Image.fromarray(imageData, "RGB")
        image.save("out/{}.png".format(uuid.uuid4()))
    except Exception as e:
        print(e)
