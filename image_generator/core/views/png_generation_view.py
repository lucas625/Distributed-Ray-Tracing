"""View for generating the png image."""

from django.http import FileResponse
from rest_framework.views import APIView

from core.business import ImageGeneratorBusiness


class PngGenerationView(APIView):
    """View for generating the png image."""

    http_method_names = ['post']

    def post(self, request, *args, **kwargs):
        """
        Receives the post request for generating the png image.
        """
        color_matrix = request.data.get("Colors")
        generated_image = ImageGeneratorBusiness.generate_png_image(color_matrix)

        return FileResponse(generated_image)
