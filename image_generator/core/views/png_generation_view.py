"""View for generating the png image."""

import os
import shutil
from tempfile import NamedTemporaryFile
import uuid

from django.conf import settings
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

        temporary_images_path = os.path.join(settings.MEDIA_ROOT, 'core', 'temporary_images')
        if not os.path.exists(temporary_images_path):
            os.makedirs(temporary_images_path)

        temporary_file_path = os.path.join(temporary_images_path, '{}.png'.format(uuid.uuid4()))
        image.save(temporary_file_path, format='png')

        temporary_file = NamedTemporaryFile(mode='w+b', suffix='png')
        with open(temporary_file_path, 'rb') as saved_image:
            shutil.copyfileobj(saved_image, temporary_file)
        os.remove(temporary_file_path)
        temporary_file.seek(0, 0)

        return FileResponse(temporary_file)
