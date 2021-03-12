"""View for running the path tracing."""

import os

from django.http import FileResponse
from django.conf import settings
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
        path_tracing_response = requests.post(
            os.path.join(settings.RAY_TRACING_ADDRESS, 'path-tracing'), json=request.data)
        color_matrix = path_tracing_response.json()

        response = requests.post(os.path.join(settings.IMAGE_GENERATOR_ADDRESS, 'api', 'png'), json=color_matrix)

        return FileResponse(response.content)
