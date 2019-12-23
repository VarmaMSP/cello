import { Curation } from 'types/app'

export const CURATION_ADD = 'curation/add'
export const CURATION_ADD_PODCASTS = 'curation/add_podcasts'

interface AddAction {
  type: typeof CURATION_ADD
  curations: Curation[]
}

interface AddPodcastsAction {
  type: typeof CURATION_ADD_PODCASTS
  curationId: string
  podcastIds: string[]
}

export type CurationActionTypes = AddAction | AddPodcastsAction
