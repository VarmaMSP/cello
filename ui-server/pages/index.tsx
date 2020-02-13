import HomeView from 'components/home_view/home_view'
import PageLayout from 'components/page_layout'
import { IndexPageSeo } from 'components/seo'
import React from 'react'
import { connect } from 'react-redux'
import { getIsUserSignedIn } from 'selectors/session'
import { AppState } from 'store'

interface StateToProps {
  isUserSignedIn: boolean
}

const IndexPage: React.FC<StateToProps> = () => {
  return (
    <>
      <IndexPageSeo />
      <PageLayout>
        <HomeView />
      </PageLayout>
    </>
  )
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  IndexPage,
)
