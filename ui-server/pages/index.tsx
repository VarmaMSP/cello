import HomeView from 'components/home_view'
import PageLayout from 'components/page_layout'
import { IndexPageSeo } from 'components/seo'
import React from 'react'

const IndexPage: React.FC<{}> = () => {
  return (
    <>
      <IndexPageSeo />
      <PageLayout>
        <HomeView />
      </PageLayout>
    </>
  )
}

export default IndexPage
