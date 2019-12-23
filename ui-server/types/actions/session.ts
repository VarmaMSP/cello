export const SESSION_INIT = 'session/init'
export const SESSION_DELETE = 'session/delete'
export const SESSION_SUBSCRIBE_PODCASTS = 'session/subscribe_podcasts'
export const SESSION_UNSUBSCRIBE_PODCASTS = 'session/unsubscribe_podcasts'

interface InitAction {
  type: typeof SESSION_INIT
  userId: string
}

interface DeleteAction {
  type: typeof SESSION_DELETE
}

interface SubscribePodcastsAction {
  type: typeof SESSION_SUBSCRIBE_PODCASTS
  podcastIds: string[]
}

interface UnsubscribePodcastsAction {
  type: typeof SESSION_UNSUBSCRIBE_PODCASTS
  podcastIds: string[]
}

export type SessionActionTypes =
  | InitAction
  | DeleteAction
  | SubscribePodcastsAction
  | UnsubscribePodcastsAction
