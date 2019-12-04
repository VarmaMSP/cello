import { Podcast, User } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function init(): Promise<{
  user: User | undefined
  subscriptions: Podcast[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/me`,
  })

  return {
    user: data.user && unmarshal.user(data.user),
    subscriptions: (data.subscriptions || []).map(unmarshal.podcast),
  }
}

export async function signOut(): Promise<void> {
  await doFetch({
    method: 'GET',
    urlPath: `/signout`,
  })
}
