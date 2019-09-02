import { NextPageContext as Page } from 'next'
import { NextJSContext as Wrapper } from 'next-redux-wrapper'
import { AppContext as App } from 'next/app'

import { AppState } from '../store'
import { AppActions } from './actions'

export interface PageContext extends Page, Wrapper<AppState, AppActions> {}

export interface AppContext extends App {
  ctx: PageContext
}
