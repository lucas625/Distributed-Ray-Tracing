"""Business object for running the path tracing."""

import os

from django.conf import settings
import requests


class PathTracingBusiness:
    """Business object for running the path tracing."""

    @staticmethod
    def run_path_tracing(path_tracing_data):
        """
        Runs the path tracing and returns the generated file.
        :param dict path_tracing_data:
        :return bytes:
        """
        path_tracing_response = requests.post(
            os.path.join(settings.RAY_TRACING_ADDRESS, 'path-tracing'), json=path_tracing_data)
        color_matrix = path_tracing_response.json()

        image_response = requests.post(os.path.join(settings.IMAGE_GENERATOR_ADDRESS, 'api', 'png'), json=color_matrix)
        return image_response.content
