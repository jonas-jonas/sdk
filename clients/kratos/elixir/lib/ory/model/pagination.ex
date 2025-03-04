# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.Pagination do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :"page",
    :"per_page"
  ]

  @type t :: %__MODULE__{
    :"page" => integer() | nil,
    :"per_page" => integer() | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.Pagination do
  def decode(value, _options) do
    value
  end
end

