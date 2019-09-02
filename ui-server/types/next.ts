import { NextPageContext } from 'next'
import { NextJSContext } from 'next-redux-wrapper'

import { AppState } from '../store'
import { AppActions } from './actions'

export type PageContext = NextPageContext & NextJSContext<AppState, AppActions>
