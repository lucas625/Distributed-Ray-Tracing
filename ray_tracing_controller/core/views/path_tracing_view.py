"""View for running the path tracing."""

from django.http import FileResponse
from rest_framework.views import APIView

from core.business import PathTracingBusiness


class PathTracingView(APIView):
    """View for running the path tracing."""

    http_method_names = ['post', 'get']

    def post(self, request, *args, **kwargs):
        """
        Receives the post request for running the path tracing.
        The request must contain JSON data as in the example objects.
        """
        image_as_bytes = PathTracingBusiness.run_path_tracing(request.data)
        return FileResponse(image_as_bytes)
