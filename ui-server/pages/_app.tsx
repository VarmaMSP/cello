import { getSignedInUser } from 'actions/user'
import AudioPlayer from 'components/audio_player'
import NavbarSide from 'components/navbar_side'
import NavbarTop from 'components/navbar_top'
import SigninModal from 'components/signin_modal'
import withRedux from 'next-redux-wrapper'
import { AppProps, Container } from 'next/app'
import Head from 'next/head'
import Router from 'next/router'
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { bindActionCreators } from 'redux'
import { makeStore } from 'store'
import * as T from 'types/actions'
import { ScreenWidth } from 'types/app'
import { AppContext, PageContext } from 'types/utilities'
import '../styles/index.css'

export default withRedux(makeStore)(
  class MyApp extends Component<AppProps & PageContext> {
    static async getInitialProps({ Component, ctx }: AppContext) {
      const { query, asPath: currentUrl, store } = ctx
      const pagePreventReload = store.getState().browser.pagePreventReload

      let scrollY = 0
      if (currentUrl === pagePreventReload.url) {
        scrollY = pagePreventReload.scrollY
      }
      if (currentUrl !== pagePreventReload.url && Component.getInitialProps) {
        await Component.getInitialProps(ctx)
      }
      return { pageProps: { ...query, scrollY } }
    }

    componentDidMount() {
      const {
        store: { dispatch, getState },
      } = this.props

      /*
       * Redux store maintains the following data regrading routing
       *  - A stack [pages] of {url, scrollPosition} pairs for all the pages
       *    that can be reached by clicking back button in the same order
       *  - A page [pagePreventReload] representing a page that user has previously
       *    navigated, and thus can be loaded from store
       */
      window.history.scrollRestoration = 'manual'

      Router.events.on('routeChangeStart', (toUrl) => {
        const pagePreventReload = getState().browser.pagePreventReload
        // Preventing push_page when user is going to previous page
        if (pagePreventReload.url !== toUrl) {
          // Push page from which user clicked on back
          dispatch({
            type: T.PUSH_PAGE,
            page: { url: Router.asPath, scrollY: window.scrollY },
          })
        }
      })

      Router.beforePopState(({ as: toUrl }) => {
        const pages = getState().browser.pages
        // Pop State will be called when users clicks on either back or next button
        // Preventing pop_state when user is going to next page
        if (pages.length > 0 && pages[0].url === toUrl) {
          // Set previous page to prevent load
          dispatch({ type: T.SET_PAGE_PREVENT_RELOAD, page: pages[0] })
          dispatch({ type: T.POP_PAGE })
        }
        return true
      })

      /*
       * Listen to screen width changes
       */
      this.handleScreenWidthChange()
      window.addEventListener('resize', this.handleScreenWidthChange)

      /*
       * Listen to route changes
       */
      Router.events.on('routeChangeComplete', (toUrl) => {
        dispatch({ type: T.SET_CURRENT_PATH_NAME, pathName: toUrl })
      })

      /*
       ** Try to get signed in user session details
       */
      bindActionCreators(getSignedInUser, dispatch)()
    }

    handleScreenWidthChange = () => {
      const setScreenWidth = (s: ScreenWidth) =>
        this.props.store.dispatch({ type: T.SET_SCREEN_WIDTH, width: s })

      const screenWidth = window.innerWidth
      if (screenWidth >= 1024) {
        return setScreenWidth('LG')
      } else if (screenWidth >= 768) {
        return setScreenWidth('MD')
      } else {
        return setScreenWidth('SM')
      }
    }

    render() {
      const { Component, pageProps, store } = this.props
      return (
        <Container>
          <Head>
            <title>phenopod</title>
          </Head>

          {/* Order components by z-axis */}
          <Provider store={store}>
            {/* base padding */}
            <div className="pl-4 pr-4 pt-20 pb-64 z-0">
              {/* additonal padding for large screens */}
              <div className="lg:pl-60 lg:pr-1 lg:pb-36">
                {/* additonal padding for extra large screens */}
                <div className="xl:pl-20 xl:pr-40">
                  <Component {...pageProps} />
                </div>
              </div>
            </div>
            <NavbarTop />
            <AudioPlayer />
            <NavbarSide />
            <SigninModal />
          </Provider>
        </Container>
      )
    }
  },
)
