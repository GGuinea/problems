defmodule BoutiqueInventory do
  def sort_by_price([_ | _] = inventory) do
    Enum.sort_by(inventory, & &1.price, :asc)
  end

  def sort_by_price([]), do: []

  def with_missing_price([_ | _] = inventory) do
    Enum.filter(inventory, &(&1.price == nil))
  end

  def with_missing_price([]), do: []

  def update_names([_ | _] = inventory, old_word, new_word) do
    Enum.map(inventory, fn item = %{name: name} ->
      %{item | name: String.replace(name, old_word, new_word)}
    end)
  end

  def update_names([], _old_word, _new_word), do: []

  def increase_quantity(item, count) do
    %{item | quantity_by_size: Map.new(item.quantity_by_size, fn {k, v} -> {k, v + count} end)}
  end

  def total_quantity(item) do
    Enum.reduce(item.quantity_by_size, 0, fn {_size, count}, acc ->
      acc + count
    end)
  end
end
