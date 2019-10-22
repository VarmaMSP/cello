import { getCurrentUser } from 'actions/user'
import AudioPlayer from 'components/audio_player'
import Modals from 'components/modals'
import NavbarSide from 'components/navbar_side'
import NavbarTop from 'components/navbar_top'
import withRedux from 'next-redux-wrapper'
import { DefaultSeo } from 'next-seo'
import { AppProps, Container } from 'next/app'
import Router from 'next/router'
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { bindActionCreators } from 'redux'
import { makeStore } from 'store'
import * as T from 'types/actions'
import { ViewportSize } from 'types/app'
import { AppContext, PageContext } from 'types/utilities'
import '../styles/index.css'

export default withRedux(makeStore)(
  class MyApp extends Component<AppProps & PageContext> {
    static async getInitialProps({ Component, ctx }: AppContext) {
      const { query, asPath: currentUrlPath, store } = ctx
      const prevPage = store.getState().browser.previousPage.page

      let scrollY = 0
      if (currentUrlPath === prevPage.urlPath) {
        scrollY = prevPage.scrollY
      }
      if (currentUrlPath !== prevPage.urlPath && Component.getInitialProps) {
        await Component.getInitialProps(ctx)
      }
      return { pageProps: { ...query, scrollY } }
    }

    componentDidMount() {
      const {
        store: { dispatch, getState },
      } = this.props

      /*
       * previous_page field in redux stor is used to store previous pages
       *  - `stack` contains all the pages that can be reached by clicking back button
       *    , in the same order
       *  - When user clicks back the `stack` is poped and the result is store in
       *    `page` field.
       *  - When a page loads, it can compare its urlPath to `previous_page.page` and determine
       *    wheather to set scroll, load data from store etc.
       */
      window.history.scrollRestoration = 'manual'

      Router.events.on('routeChangeStart', (urlPath) => {
        const prevPage = getState().browser.previousPage.page
        // Preventing push when user is going to previous page
        if (prevPage.urlPath !== urlPath) {
          dispatch({
            type: T.PUSH_PREVIOUS_PAGE_STACK,
            page: { urlPath: Router.asPath, scrollY: window.scrollY }, // Push page from which user clicked on back
          })
        }
      })

      Router.beforePopState(({ as: toUrlPath }) => {
        const state = getState()

        // Close modal id opened and prevent route change
        if (state.ui.showModal.type !== 'NONE') {
          dispatch({ type: T.CLOSE_MODAL })
          return false
        }

        const stack = state.browser.previousPage.stack
        // Preventing pop_state when user is going to next page
        if (stack.length > 0 && stack[0].urlPath === toUrlPath) {
          dispatch({ type: T.SET_PREVIOUS_PAGE, page: stack[0] })
          dispatch({ type: T.POP_PREVIOUS_PAGE_STACK })
        }
        return true
      })

      /*
       * Listen to screen width changes
       */
      this.handleViewportSizeChange()
      window.addEventListener('resize', this.handleViewportSizeChange)

      /*
       * Listen to route changes
       */
      Router.events.on('routeChangeComplete', (toUrlPath) => {
        dispatch({ type: T.SET_CURRENT_URL_PATH, urlPath: toUrlPath })
      })

      /*
       * Try to get signed in user session details
       */
      bindActionCreators(getCurrentUser, dispatch)()
    }

    handleViewportSizeChange = () => {
      const setViewportSize = (s: ViewportSize) =>
        this.props.store.dispatch({ type: T.SET_VIEWPORT_SIZE, size: s })

      const width = window.innerWidth
      if (width >= 1024) {
        return setViewportSize('LG')
      } else if (width >= 768) {
        return setViewportSize('MD')
      } else {
        return setViewportSize('SM')
      }
    }

    render() {
      const { Component, pageProps, store } = this.props
      return (
        <Container>
          {/* Default seo that can be overidden by individual pages */}
          <DefaultSeo
            openGraph={{
              site_name: 'Phenopod',
            }}
            twitter={{
              handle: '@phenopod',
              site: '@phenopod',
              cardType: 'summary',
            }}
            facebook={{
              appId: '526472207897979',
            }}
          />

          {/* Order components by z-axis */}
          <Provider store={store}>
            {/* base padding */}
            <div className="pl-4 pr-4 pt-20 pb-40 z-0">
              {/* additonal padding for large screens */}
              <div className="lg:pl-60 lg:pr-1 lg:pb-28">
                {/* additonal padding for extra large screens */}
                <div className="xl:pl-20 xl:pr-40">
                  <Component {...pageProps} />
                </div>
              </div>
            </div>
            <NavbarTop />
            <AudioPlayer />
            <NavbarSide />
            <Modals />
          </Provider>
        </Container>
      )
    }
  },
)
