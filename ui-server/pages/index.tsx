import About from 'components/about'
import { iconMap } from 'components/icon'
import PageLayout from 'components/page_layout'
import SignInButton from 'components/sign_in_button'
import React from 'react'
import { connect } from 'react-redux'
import { getIsUserSignedIn } from 'selectors/session'
import { AppState } from 'store'
import { getAssetUrl } from 'utils/dom'

interface StateToProps {
  isUserSignedIn: boolean
}

const IndexPage: React.FC<StateToProps> = ({ isUserSignedIn }) => {
  const LogoIcon = iconMap['logo-lg']

  return (
    <PageLayout>
      <div className="flex md:flex-row-reverse flex-col md:-mb-32">
        <img
          src={getAssetUrl('home')}
          className="md:flex-1 md:w-5/12 md:-mt-8 -mt-12 -mr-6 "
        />
        <div className="flex-none md:w-1/2 md:mt-0 -mt-10 md:pt-16 ">
          <LogoIcon className="mb-2" />
          <h3 className="mb-8 tracking-wide text-gray-900 text-lg">
            {'The Best Online Podcast Player'}
          </h3>
          <p className="text-sm tracking-wide leading-loose">{'Browse from'}</p>
          <p className="mb-2 tracking-wide">
            <span className="text-xl text-teal-700 font-bold">
              {'638,417 Podcasts'}
            </span>
            <span className="text-sm mx-1">{' and '}</span>
            <span className="text-xl text-teal-700 font-bold">
              {'21,708,559 Episodes'}
            </span>
          </p>
          <p className="text-sm tracking-wide leading-loose">
            {'And search for'}
          </p>
          <p className="mb-8 tracking-wide text-xl text-teal-700 font-bold">
            {'Topic, Person, Podcast, Episode, Publisher...'}
          </p>
          {!isUserSignedIn && (
            <>
              <p className="mb-5 tracking-wide break-words">
                {
                  'Keep track of your subscriptions, listen history and curate episodes playlists. Sign Up / Sign In to get started.'
                }
              </p>
              <div className="h-10 w-48">
                <SignInButton />
              </div>
            </>
          )}
          <hr className="md:hidden mt-10 mb-6" />
          <div className="md:hidden text-center">
            <About />
          </div>
        </div>
      </div>
    </PageLayout>
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
