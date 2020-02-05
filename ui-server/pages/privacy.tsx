import PageLayout from 'components/page_layout'
import { PrivacyPageSeo } from 'components/seo'
import React from 'react'

export default class PrivacyPage extends React.Component<{}> {
  render() {
    return (
      <>
        <PrivacyPageSeo />
        <PageLayout>
          <div>
            <h1 className="pt-3 pb-1 text-xl tracking-wide text-gray-800 font-medium">
              {'Privacy'}
            </h1>
            <hr className="mb-3" />
            <div className="text-gray-700 tracking-wide font-normal">
              <p className="pb-2">
                {
                  'This site uses cookies to save your preferences and playback progress'
                }
                <br />
                <br />
                {
                  'We collect your basic information (name, username and email) when you sign in with social account,'
                }
                <br />{' '}
                {'and We will not share those details with any third party.'}
              </p>
            </div>
          </div>
        </PageLayout>
      </>
    )
  }
}
