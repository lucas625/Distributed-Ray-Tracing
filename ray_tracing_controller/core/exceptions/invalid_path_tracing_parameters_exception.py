# !/usr/bin/env python3
# -*- coding: utf-8 -*-

"""Raised when the path tracing parameters are bigger than they should be."""


class InvalidPathTracingParametersException(Exception):
    """Raised when the path tracing parameters are bigger than they should be."""

    def __init__(self, max_path_tracing_parameters, received_path_tracing_parameters):
        """
        :param dict max_path_tracing_parameters:
        :param dict received_path_tracing_parameters:
        """
        super().__init__(
            'Invalid path tracing parameters, please use a more reasonable approach.'
            '\nReceived Parameters: width: {}, height: {}, rays per pixel: {} recursions: {}.'
            '\n Max Parameters: width: {}, height: {}, rays per pixel: {} recursions: {}'.format(
                received_path_tracing_parameters.get('width'), received_path_tracing_parameters.get('height'),
                received_path_tracing_parameters.get('raysPerPixel'),
                received_path_tracing_parameters.get('recursions'),
                max_path_tracing_parameters.get('width'), max_path_tracing_parameters.get('height'),
                max_path_tracing_parameters.get('raysPerPixel'), max_path_tracing_parameters.get('recursions')))
