

# SubmitSelfServiceRecoveryFlowBody


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**csrfToken** | **String** | Sending the anti-csrf token is only required for browser login flows. |  [optional] |
|**email** | **String** | Email to Recover  Needs to be set when initiating the flow. If the email is a registered recovery email, a recovery link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email |  |
|**method** | **String** | Method supports &#x60;link&#x60; only right now. |  |



