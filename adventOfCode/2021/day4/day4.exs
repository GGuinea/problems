defmodule Day4 do
  def get_input(file_name) do
    {:ok, body} = File.read(file_name)

    body
    |> String.split("\n", trim: true)
    |> Enum.map(&(&1 |> String.to_charlist() |> List.to_tuple()))
  end
end
