defmodule KitchenCalculator do
  def get_volume(volume_pair) do
    elem(volume_pair, 1)
  end

  def to_milliliter({:cup, num}) do
    {:milliliter, num * 240}
  end

  def to_milliliter({:fluid_ounce, num}) do
    {:milliliter, num * 30}
  end

  def to_milliliter({:teaspoon, num}) do
    {:milliliter, num * 5}
  end

  def to_milliliter({:tablespoon, num}) do
    {:milliliter, num * 15}
  end

  def to_milliliter({:milliliter, num}) do
    {:milliliter, num}
  end

  def from_milliliter(volume_pair, :cup) do
    {:cup, get_volume(volume_pair) / 240}
  end

  def from_milliliter(volume_pair, :fluid_ounce) do
    {:fluid_ounce, get_volume(volume_pair) / 30}
  end

  def from_milliliter(volume_pair, :teaspoon) do
    {:teaspoon, get_volume(volume_pair) / 5}
  end

  def from_milliliter(volume_pair, :tablespoon) do
    {:tablespoon, get_volume(volume_pair) / 15}
  end

  def from_milliliter(volume_pair, :milliliter) do
    volume_pair
  end

  def convert(volume_pair, unit) do
    volume_pair
    |> to_milliliter()
    |> from_milliliter(unit)
  end
end
