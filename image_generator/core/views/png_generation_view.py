"""View for generating the png image."""

import os
from io import BytesIO
import uuid

from django.core.files.uploadedfile import InMemoryUploadedFile
from django.http import FileResponse
from PIL import Image
import numpy as np
from rest_framework.views import APIView


class PngGenerationView(APIView):
    """View for generating the png image."""

    http_method_names = ['post', 'get']

    def post(self, request, *args, **kwargs):
        """
        Receives the post request for generating the png image.
        """
        color_matrix = request.data.get("Colors")
        width = len(color_matrix[0])
        height = len(color_matrix)
        image_data = np.zeros((height, width, 3), dtype=np.uint8)
        for lineIndex in range(height):
            for columnIndex in range(width):
                for colorIndex in range(3):
                    image_data[lineIndex][columnIndex][colorIndex] = color_matrix[lineIndex][columnIndex][colorIndex]
        image = Image.fromarray(image_data, "RGB")
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
