"""Bean for the color matrix."""

import threading


class ColorMatrixBean:
    """Bean for the color matrix."""

    def __init__(self, height, width):
        """
        Initializes a ColorMatrixBean.
        :param int height:
        :param int width:
        """
        self._lock = threading.Lock()

        color_matrix = [[0, 0, 0] * width] * height
        self._color_matrix = color_matrix

    @property
    def color_matrix(self):
        """
        Getter for the color matrix.
        :return list[list[list[int]]]:
        """
        return self._color_matrix

    def set_color_matrix_by_window(self, new_color_matrix, start_line, start_column, end_line, end_column):
        """
        Setter for the color matrix using a window.
        :param list[list[list[int]]] new_color_matrix:
        :param int start_line:
        :param int start_column:
        :param int end_line:
        :param int end_column:
        """
        with self._lock:
            for lineIndex in range(start_line, end_line):
                for columnIndex in range(start_column, end_column):
                    self._color_matrix[lineIndex][columnIndex] = new_color_matrix[lineIndex][columnIndex]

    def to_dto(self):
        """
        Parses the bean to a dict that can be send as json.
        :return dict:
        """
        return dict(Colors=self.color_matrix)
