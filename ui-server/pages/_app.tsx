import AudioPlayer from 'components/audio_player'
import MainContainer from 'components/main_container'
import NavbarSide from 'components/navbar_side'
import NavbarTop from 'components/navbar_top'
import SigninModal from 'components/signin_modal'
import withRedux from 'next-redux-wrapper'
import { AppProps, Container } from 'next/app'
import Head from 'next/head'
import Router from 'next/router'
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { makeStore } from 'store'
import { POP_PAGE, PUSH_PAGE, SET_PAGE_PREVENT_RELOAD } from 'types/actions'
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
      window.history.scrollRestoration = 'manual'

      /*
       ** Redux store maintains the following data regrading routing
       **  - A stack [pages] of {url, scrollPosition} pairs for all the pages
       **    that can be reached by clicking back button in the same order
       **  - A page [pagePreventReload] representing a page that user has previously
       **    navigated, and thus can be loaded from store
       */

      const { store } = this.props
      Router.events.on('routeChangeStart', (toUrl) => {
        const pagePreventReload = store.getState().browser.pagePreventReload
        // Preventing push_page when user is going to previous page
        if (pagePreventReload.url !== toUrl) {
          // Push page from which user clicked on back
          store.dispatch({
            type: PUSH_PAGE,
            page: { url: Router.asPath, scrollY: window.scrollY },
          })
        }
      })

      Router.beforePopState(({ as: toUrl }) => {
        const pages = store.getState().browser.pages
        // Pop State will be called when users clicks on either back or next button
        // Preventing pop_state when user is going to next page
        if (pages.length > 0 && pages[0].url === toUrl) {
          // Set previous page to prevent load
          store.dispatch({ type: SET_PAGE_PREVENT_RELOAD, page: pages[0] })
          store.dispatch({ type: POP_PAGE })
        }

        return true
      })
    }

    render() {
      const { Component, pageProps, store } = this.props
      return (
        <Container>
          <Head>
            <title>phenopod</title>
          </Head>

          <Provider store={store}>
            {/* Order components by z-axis */}
            <MainContainer>
              <Component {...pageProps} />
            </MainContainer>
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
