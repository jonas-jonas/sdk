# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody do
  @moduledoc """
  SubmitSelfServiceRegistrationFlowWithPasswordMethodBody is used to decode the registration form payload when using the password method.
  """

  @derive [Poison.Encoder]
  defstruct [
    :csrf_token,
    :method,
    :password,
    :traits
  ]

  @type t :: %__MODULE__{
    :csrf_token => String.t | nil,
    :method => String.t,
    :password => String.t,
    :traits => map()
  }
end

defimpl Poison.Decoder, for: Ory.Model.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody do
  def decode(value, _options) do
    value
  end
end

