import { PodcastActionTypes } from './podcast'
import { PlayerActionTypes } from './player'

export type AppActions = PodcastActionTypes | PlayerActionTypes

export * from './podcast'
export * from './player'
