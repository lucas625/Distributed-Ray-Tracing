"""Business object for running the path tracing."""

import os
import threading

from django.conf import settings
import requests

from core.beans import ColorMatrixBean


class PathTracingBusiness:
    """Business object for running the path tracing."""

    @staticmethod
    def run_path_tracing(path_tracing_data):
        """
        Runs the path tracing and returns the generated file.
        :param dict path_tracing_data:
        :return bytes:
        """
        height = path_tracing_data.get('pixelScreen').get('height')
        width = path_tracing_data.get('pixelScreen').get('width')
        color_matrix = ColorMatrixBean(height, width)
        thread_count = 0
        max_thread_count = 3
        threads = []
        window_size = 50

        for line_index in range(height):
            if thread_count < max_thread_count:
                end_line = line_index+window_size if line_index + window_size <= height else height
                arguments = (path_tracing_data, color_matrix, line_index, end_line, width)
                line_index += window_size
                thread = threading.Thread(target=PathTracingBusiness._perform_windowed_path_tracing, args=arguments)
                threads.append(thread)
                thread.start()
            else:
                for thread in threads:
                    thread.join()
                threads = []
        for thread in threads:
            thread.join()
        image_response = requests.post(
            os.path.join(settings.IMAGE_GENERATOR_ADDRESS, 'api', 'png'), json=color_matrix.to_dto())
        return image_response.content

    @staticmethod
    def _perform_windowed_path_tracing(path_tracing_data, color_matrix, start_line, end_line, width):
        """
        :param dict path_tracing_data:
        :param ColorMatrixBean color_matrix:
        :param int start_line:
        :param int end_line:
        :param int width:
        """
        original_path_tracing_parameters = path_tracing_data.get('pathTracingParameters')
        copied_data = dict(
            pathTracingParameters=dict(
                raysPerPixel=original_path_tracing_parameters.get('raysPerPixel'),
                recursions=original_path_tracing_parameters.get('recursions'),
                windowStartLine=start_line,
                windowStartColumn=0,
                windowEndLine=end_line,
                windowEndColumn=width
            ),
            objects=path_tracing_data.get('objects'),
            pixelScreen=path_tracing_data.get('pixelScreen'),
            sceneCamera=path_tracing_data.get('sceneCamera'),
            lights=path_tracing_data.get('lights')
        )
        path_tracing_response = requests.post(
            os.path.join(settings.RAY_TRACING_ADDRESS, 'path-tracing'), json=copied_data)
        response_color_matrix = path_tracing_response.json().get('Colors')
        color_matrix.set_color_matrix_by_window(response_color_matrix, start_line, 0, end_line, width)

