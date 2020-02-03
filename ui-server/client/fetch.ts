import fetch from 'isomorphic-unfetch'
import { ApiResponse } from 'types/models/api_response'

export interface FetchRequest {
  method: 'GET' | 'PUT' | 'POST'
  urlPath: string
  body?: object
}

export interface FetchResponse {
  responseHeaders: { [key: string]: string }
  data: any
}

export interface FetchException {
  url: string
  statusCode: number
  responseHeaders: { [key: string]: string }
  err?: string
}

export async function doFetch({
  method,
  urlPath,
  body,
}: FetchRequest): Promise<FetchResponse> {
  // Make Request
  let url = `${getBaseUrl()}${urlPath}`
  let response: Response
  try {
    response = await fetch(url, {
      method,
      body: body ? JSON.stringify(body) : undefined,
      credentials: 'include',
    })
  } catch (err) {
    throw makeFetchException(url, response!.status, {}, err.toString())
  }

  // Check for non 2xx { not OK }
  if (response!.status.toString()[0] !== '2') {
    throw makeFetchException(url, response!.status, {})
  }

  // Parse body
  let data: object = {}
  try {
    if (response!.headers.get('Content-Type') === 'application/json') {
      data = await response.json()
    }
  } catch (err) {
    const errMsg = 'JSON: error parsing body'
    throw makeFetchException(url, response!.status, {}, errMsg)
  }

  return <FetchResponse>{
    responseHeaders: {
      Location: response!.headers.get('Location') || '',
    },
    data,
  }
}

export async function doFetch_({
  method,
  urlPath,
  body,
}: FetchRequest): Promise<ApiResponse> {
  // Make Request
  let url = `${getBaseUrl()}${urlPath}`
  let response: Response
  try {
    response = await fetch(url, {
      method,
      body: body ? JSON.stringify(body) : undefined,
      credentials: 'include',
    })
  } catch (err) {
    throw makeFetchException(url, response!.status, {}, err.toString())
  }

  // Check for non 2xx { not OK }
  if (response!.status.toString()[0] !== '2') {
    throw makeFetchException(url, response!.status, {})
  }

  // Parse body
  let data: object = {}
  try {
    if (response!.headers.get('Content-Type') === 'application/json') {
      data = await response.json()
    }
  } catch (err) {
    const errMsg = 'JSON: error parsing body'
    throw makeFetchException(url, response!.status, {}, errMsg)
  }

  return new ApiResponse(data)
}

function getBaseUrl(): string {
  if (process.browser) {
    return process.env.NODE_ENV === 'production'
      ? 'https://phenopod.com/api'
      : 'http://localhost:8080/api'
  }
  return 'http://localhost:8081'
}

function makeFetchException(
  url: string,
  statusCode: number,
  responseHeaders: { [key: string]: string },
  err?: string,
) {
  return <FetchException>{ url, statusCode, responseHeaders, err }
}
