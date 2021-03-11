"""View for running the path tracing."""

import os
from io import BytesIO
import uuid

from django.conf import settings
from django.core.files.uploadedfile import InMemoryUploadedFile
from django.http import FileResponse
from PIL import Image
import numpy as np
import requests
from rest_framework.views import APIView


class PathTracingView(APIView):
    """View for running the path tracing."""

    http_method_names = ['post', 'get']

    def post(self, request, *args, **kwargs):
        """
        Receives the post request for running the path tracing.
        The request must contain JSON data as in the example objects.
        """
        response = requests.post(os.path.join(settings.RAY_TRACING_ADDRESS, 'path-tracing'), json=request.data)

        colorMatrix = response.json()["Colors"]
        width = len(colorMatrix[0])
        height = len(colorMatrix)
        imageData = np.zeros((height, width, 3), dtype=np.uint8)
        for lineIndex in range(height):
            for columnIndex in range(width):
                for colorIndex in range(3):
                    imageData[lineIndex][columnIndex][colorIndex] = colorMatrix[lineIndex][columnIndex][colorIndex]
        image = Image.fromarray(imageData, "RGB")
        image_as_bytes = BytesIO()
        image.save(image_as_bytes, format='png')

        in_memory_image = InMemoryUploadedFile(
            file=image_as_bytes,
            field_name='test_file_{}.png'.format(uuid.uuid4()),
            name='test_file_{}.png'.format(uuid.uuid4()),
            content_type='image/png',
            size=None,
            charset=None)

        return FileResponse(in_memory_image)
