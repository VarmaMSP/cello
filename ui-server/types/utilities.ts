import { NextPageContext as Page } from 'next'
import { NextJSContext as Wrapper } from 'next-redux-wrapper'
import { AppContext as App } from 'next/app'
import { AppState } from '../store'
import { AppActions } from './actions'
import { Entity } from './app'

//
// Extend NextJs Context interfaces with Context Pages in by next redux wrapper
//
export interface PageContext extends Page, Wrapper<AppState, AppActions> {}

export interface AppContext extends App {
  ctx: PageContext
}

//
// types with phantom inputs to better type selectors and keep my sanity
//
export type $Id<_E extends Entity> = string

export type MapById<E extends Entity> = { [id: string]: E }

export type MapOneToOne<_E1 extends Entity, E2 extends Entity> = {
  [id: string]: E2
}

export type MapOneToMany<_E1 extends Entity, _E2 extends Entity> = {
  [id: string]: string[]
}
