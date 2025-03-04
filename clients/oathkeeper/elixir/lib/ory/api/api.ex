# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Api.Api do
  @moduledoc """
  API calls for all endpoints tagged `Api`.
  """

  alias Ory.Connection
  import Ory.RequestBuilder


  @doc """
  Access Control Decision API
  > This endpoint works with all HTTP Methods (GET, POST, PUT, ...) and matches every path prefixed with /decision.  This endpoint mirrors the proxy capability of ORY Oathkeeper's proxy functionality but instead of forwarding the request to the upstream server, returns 200 (request should be allowed), 401 (unauthorized), or 403 (forbidden) status codes. This endpoint can be used to integrate with other API Proxies like Ambassador, Kong, Envoy, and many more.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, nil} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec decisions(Tesla.Env.client, keyword()) :: {:ok, nil} | {:ok, Ory.Model.GetWellKnownJsonWebKeys500Response.t} | {:error, Tesla.Env.t}
  def decisions(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/decisions")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, false},
      { 401, %Ory.Model.GetWellKnownJsonWebKeys500Response{}},
      { 403, %Ory.Model.GetWellKnownJsonWebKeys500Response{}},
      { 404, %Ory.Model.GetWellKnownJsonWebKeys500Response{}},
      { 500, %Ory.Model.GetWellKnownJsonWebKeys500Response{}}
    ])
  end

  @doc """
  Retrieve a rule
  Use this method to retrieve a rule from the storage. If it does not exist you will receive a 404 error.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - id (String.t): 
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, Ory.Model.Rule.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec get_rule(Tesla.Env.client, String.t, keyword()) :: {:ok, Ory.Model.Rule.t} | {:ok, Ory.Model.GetWellKnownJsonWebKeys500Response.t} | {:error, Tesla.Env.t}
  def get_rule(connection, id, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/rules/#{id}")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %Ory.Model.Rule{}},
      { 404, %Ory.Model.GetWellKnownJsonWebKeys500Response{}},
      { 500, %Ory.Model.GetWellKnownJsonWebKeys500Response{}}
    ])
  end

  @doc """
  Get service version
  This endpoint returns the service version typically notated using semantic versioning.  If the service supports TLS Edge Termination, this endpoint does not require the `X-Forwarded-Proto` header to be set.  Be aware that if you are running multiple nodes of this service, the health status will never refer to the cluster state, only to a single instance.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, Ory.Model.Version.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec get_version(Tesla.Env.client, keyword()) :: {:ok, Ory.Model.Version.t} | {:error, Tesla.Env.t}
  def get_version(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/version")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %Ory.Model.Version{}}
    ])
  end

  @doc """
  Lists cryptographic keys
  This endpoint returns cryptographic keys that are required to, for example, verify signatures of ID Tokens.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, Ory.Model.JsonWebKeySet.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec get_well_known_json_web_keys(Tesla.Env.client, keyword()) :: {:ok, Ory.Model.JsonWebKeySet.t} | {:ok, Ory.Model.GetWellKnownJsonWebKeys500Response.t} | {:error, Tesla.Env.t}
  def get_well_known_json_web_keys(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/.well-known/jwks.json")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %Ory.Model.JsonWebKeySet{}},
      { 500, %Ory.Model.GetWellKnownJsonWebKeys500Response{}}
    ])
  end

  @doc """
  Check alive status
  This endpoint returns a 200 status code when the HTTP server is up running. This status does currently not include checks whether the database connection is working.  If the service supports TLS Edge Termination, this endpoint does not require the `X-Forwarded-Proto` header to be set.  Be aware that if you are running multiple nodes of this service, the health status will never refer to the cluster state, only to a single instance.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, Ory.Model.HealthStatus.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec is_instance_alive(Tesla.Env.client, keyword()) :: {:ok, Ory.Model.GetWellKnownJsonWebKeys500Response.t} | {:ok, Ory.Model.HealthStatus.t} | {:error, Tesla.Env.t}
  def is_instance_alive(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/health/alive")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %Ory.Model.HealthStatus{}},
      { 500, %Ory.Model.GetWellKnownJsonWebKeys500Response{}}
    ])
  end

  @doc """
  Check readiness status
  This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g. the database) are responsive as well.  If the service supports TLS Edge Termination, this endpoint does not require the `X-Forwarded-Proto` header to be set.  Be aware that if you are running multiple nodes of this service, the health status will never refer to the cluster state, only to a single instance.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, Ory.Model.HealthStatus.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec is_instance_ready(Tesla.Env.client, keyword()) :: {:ok, Ory.Model.HealthNotReadyStatus.t} | {:ok, Ory.Model.HealthStatus.t} | {:error, Tesla.Env.t}
  def is_instance_ready(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/health/ready")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %Ory.Model.HealthStatus{}},
      { 503, %Ory.Model.HealthNotReadyStatus{}}
    ])
  end

  @doc """
  List all rules
  This method returns an array of all rules that are stored in the backend. This is useful if you want to get a full view of what rules you have currently in place.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
    - :limit (integer()): The maximum amount of rules returned.
    - :offset (integer()): The offset from where to start looking.
  ## Returns

  {:ok, [%Rule{}, ...]} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec list_rules(Tesla.Env.client, keyword()) :: {:ok, list(Ory.Model.Rule.t)} | {:ok, Ory.Model.GetWellKnownJsonWebKeys500Response.t} | {:error, Tesla.Env.t}
  def list_rules(connection, opts \\ []) do
    optional_params = %{
      :limit => :query,
      :offset => :query
    }
    %{}
    |> method(:get)
    |> url("/rules")
    |> add_optional_params(optional_params, opts)
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, [%Ory.Model.Rule{}]},
      { 500, %Ory.Model.GetWellKnownJsonWebKeys500Response{}}
    ])
  end
end
