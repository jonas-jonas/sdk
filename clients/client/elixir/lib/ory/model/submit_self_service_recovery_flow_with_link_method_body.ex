# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.SubmitSelfServiceRecoveryFlowWithLinkMethodBody do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :csrf_token,
    :email,
    :method
  ]

  @type t :: %__MODULE__{
    :csrf_token => String.t | nil,
    :email => String.t,
    :method => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Model.SubmitSelfServiceRecoveryFlowWithLinkMethodBody do
  def decode(value, _options) do
    value
  end
end

