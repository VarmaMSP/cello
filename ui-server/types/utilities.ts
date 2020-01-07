import { NextComponentType, NextPageContext as Page } from 'next'
import { NextJSContext as Wrapper } from 'next-redux-wrapper'
import { AppContext as App } from 'next/app'
import { AppState } from 'store'
import { AppActions } from './actions'

//
// Extend NextJs Context interfaces with Context Pages in by next redux wrapper
//
export interface PageContext extends Page, Wrapper<AppState, AppActions> {}

export interface AppContext extends App {
  ctx: PageContext
  Component: NextComponentType<PageContext> & {
    loadPropsIntoStore?: (context: PageContext) => void
  }
}

//
// types with phantom inputs to better type selectors and keep my sanity
//
export type $Id<_E> = string

export type MapById<E> = { [id: string]: E }

export type MapOneToOne<_E1, E2> = { [id: string]: E2 }

export type MapOneToMany<_E1, E2> = { [id: string]: $Id<E2>[] }
