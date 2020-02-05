import { getHomePageData } from 'actions/home'
import HomeView from 'components/home_view/home_view'
import PageLayout from 'components/page_layout'
import { ExplorePageSeo } from 'components/seo'
import React from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

interface OwnProps {
  scrollY: number
}

export default class IndexPage extends React.Component<OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    await bindActionCreators(getHomePageData, ctx.store.dispatch)()
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <>
        <ExplorePageSeo />
        <PageLayout>
          <HomeView />
        </PageLayout>
      </>
    )
  }
}
