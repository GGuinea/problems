defmodule RPG do
  defmodule Character do
    defstruct health: 100, mana: 0
  end

  defmodule LoafOfBread do
    defstruct []
  end

  defmodule ManaPotion do
    defstruct strength: 10
  end

  defmodule Poison do
    defstruct []
  end

  defmodule EmptyBottle do
    defstruct []
  end

  # Add code to define the protocol and its implementations below here...
  defprotocol Edible do
    def eat(item, character)
  end

  defimpl RPG.Edible, for: RPG.LoafOfBread do
    def eat(_item, character) do
      {nil,
       %RPG.Character{
         health: Map.get(character, :health, 0) + 5,
         mana: Map.get(character, :mana, 0)
       }}
    end
  end

  defimpl RPG.Edible, for: RPG.ManaPotion do
    def eat(item, character) do
      {%RPG.EmptyBottle{},
       %RPG.Character{
         health: Map.get(character, :health, 0),
         mana: Map.get(character, :mana, 0) + Map.get(item, :strength, 0)
       }}
    end
  end

  defimpl RPG.Edible, for: RPG.Poison do
    def eat(item, character) do
      {%RPG.EmptyBottle{},
       %RPG.Character{
         health: 0,
         mana: Map.get(character, :mana, 0) + Map.get(item, :strength, 0)
       }}
    end
  end
end
