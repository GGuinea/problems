defmodule BasketballWebsite do
  def extract_from_path(data, path) do
    paths = String.split(path, ".")
    do_extract(data, paths)
  end

  def get_in_path(data, path) do
    paths = String.split(path, ".")
    Kernel.get_in(data, paths)
  end

  defp do_extract(nil, _), do: nil

  defp do_extract(data, []), do: data

  defp do_extract(data, [head | tail]), do: do_extract(data[head], tail)
end
