"""Business object for running the path tracing."""

import os
import uuid

from django.core.files.base import ContentFile
from django.conf import settings
import requests

from core.exceptions import InvalidPathTracingParametersException

_MAX_PATH_TRACING_PARAMETERS = dict(
    width=1080,
    height=768,
    raysPerPixel=400,
    recursions=5
)


class PathTracingBusiness:
    """Business object for running the path tracing."""

    @staticmethod
    def run_path_tracing(path_tracing_data):
        """
        Runs the path tracing and returns the generated file.
        :param dict path_tracing_data:
        :return bytes:
        """
        height = int(path_tracing_data.get('pixelScreen').get('height'))
        width = int(path_tracing_data.get('pixelScreen').get('width'))
        rays_per_pixel = int(path_tracing_data.get('pathTracingParameters').get('raysPerPixel'))
        recursions = int(path_tracing_data.get('pathTracingParameters').get('recursions'))

        has_invalid_path_tracing_parameters = \
            rays_per_pixel > _MAX_PATH_TRACING_PARAMETERS.get('raysPerPixel')\
            or recursions > _MAX_PATH_TRACING_PARAMETERS.get('recursions')\
            or width > _MAX_PATH_TRACING_PARAMETERS.get('width')\
            or height > _MAX_PATH_TRACING_PARAMETERS.get('height')
        if has_invalid_path_tracing_parameters:
            raise InvalidPathTracingParametersException(_MAX_PATH_TRACING_PARAMETERS, path_tracing_data)
 
        path_tracing_data['pathTracingParameters']['windowStartLine'] = 0
        path_tracing_data['pathTracingParameters']['windowStartColumn'] = 0
        path_tracing_data['pathTracingParameters']['windowEndLine'] = height
        path_tracing_data['pathTracingParameters']['windowEndColumn'] = width

        path_tracing_response = requests.post(
            os.path.join(settings.RAY_TRACING_ADDRESS, 'path-tracing'), json=path_tracing_data)
        color_matrix = path_tracing_response.json()

        image_response = requests.post(
            os.path.join(settings.IMAGE_GENERATOR_ADDRESS, 'api', 'png'), json=color_matrix)
        return ContentFile(image_response.content, '{}-w{}-h{}-rp{}-r{}.png'.format(uuid.uuid4, width, height, rays_per_pixel, recursions))
