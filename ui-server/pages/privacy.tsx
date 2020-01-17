import React from 'react'

export default class PrivacyPage extends React.Component<{}> {
  render() {
    return (
      <>
        <h1 className="text-xl text-gray-900 pb-4">{'Privacy'}</h1>
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
            <br /> {'and We will not share those details with any third party.'}
          </p>
        </div>
      </>
    )
  }
}
