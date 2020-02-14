import classnames from 'classnames'
import Grid from 'components/grid'
import { PodcastLink } from 'components/link'
import SignInButton from 'components/sign_in_button'
import React from 'react'
import { getAssetUrl, getImageUrl } from 'utils/dom'

const acclaimedPodcasts = [
  {
    id: 'e36k4d',
    urlParam: 'making-sense-with-sam-harris-e36k4d',
    title: 'Making Sense with Sam Harris',
    author: 'Sam Harris',
  },
  {
    id: 'eXDv5e',
    urlParam: 'the-joe-rogan-experience-eXDv5e',
    title: 'The Joe Rogan Experience',
    author: 'Joe Rogan',
  },
  {
    id: 'aKLBzb',
    urlParam: 'the-mysterious-mr-epstein-aKLBzb',
    title: 'The Mysterious Mr. Epstein',
    author: 'Wondery',
  },
  {
    id: 'epgPXd',
    urlParam: 'blood-ties-epgPXd',
    title: 'Blood Ties',
    author: 'Wondery',
  },
  {
    id: 'aKLBzb',
    urlParam: 'the-mysterious-mr-epstein-aKLBzb',
    title: 'The Mysterious Mr. Epstein',
    author: 'Wondery',
  },
  {
    id: 'epgPXd',
    urlParam: 'blood-ties-epgPXd',
    title: 'Blood Ties',
    author: 'Wondery',
  },
]

export interface StateToProps {
  isUserSignedIn: boolean
}

const HomeView: React.FC<StateToProps> = ({ isUserSignedIn }) => {
  return (
    <div>
      <div className="flex md:flex-row flex-col pt-6">
        <div className="flex-1 md:pt-16">
          <h1 className="text-5xl text-center text-purple-700 font-semibold tracking-wide">
            {'Phenopod'}
          </h1>
          <h2 className="text-lg text-center text-gray-900 font-medium leading-relaxed tracking-wide">
            {'Simple, Yet Powerful Podcast Player'}
          </h2>
          <h3
            className={classnames(
              'w-3/4 mt-6 mx-auto text-sm text-center text-gray-800 font-medium tracking-wide',
              { hidden: isUserSignedIn },
            )}
          >
            {
              'Discover, Subscribe, Curate Your Favourite Podcasts And Episodes. Sign In Now To Get Started.'
            }
          </h3>
          <div
            className={classnames('w-56 h-8 mx-auto mt-3', {
              hidden: isUserSignedIn,
            })}
          >
            <SignInButton />
          </div>
        </div>
        <img
          src={getAssetUrl('listen-to-podcasts.svg')}
          className="md:block hidden w-1/2"
        />
      </div>

      <div className="flex md:flex-row flex-col mt-12 md:pl-6">
        <img
          src={getAssetUrl('powerful-search.png')}
          className="md:w-1/3 w-3/5 md:mx-0 mx-auto"
        />
        <div className="flex-1 md:pt-14 pt-6">
          <h1 className="text-2xl text-center text-black font-medium leading-relaxed tracking-wide">
            {'Powerful Search Engine'}
          </h1>
          <h2 className="md:w-3/4 w-full mt-3 mx-auto text-lg text-center text-gray-800 font-medium leading-relaxed tracking-wide">
            {
              'Search For Person, Topic, Place, Podcast, Episode.. From A Directory Of 638,417 Podcasts And 21,07,559 Episodes.'
            }
          </h2>
        </div>
      </div>

      <div className="flex md:flex-row flex-col mt-16 md:pl-6">
        <img
          src={getAssetUrl('subscriptions-feed.png')}
          className="md:w-1/3 w-3/5 md:mx-0 mx-auto"
        />
        <div className="flex-1 md:pt-12 pt-6">
          <h1 className="text-2xl text-center text-black font-medium leading-relaxed tracking-wide">
            {'Subscriptions'}
          </h1>
          <h2 className="md:w-3/4 w-full mt-3 mx-auto text-lg text-center text-gray-800 font-medium leading-relaxed tracking-wide">
            {
              'Subscribe To Your Favourite Podcasts And Get Lastest Episodes In Your Subscriptions Feed.'
            }
          </h2>
        </div>
      </div>

      <div className="flex md:flex-row flex-col mt-16 md:pl-6">
        <img
          src={getAssetUrl('sync-progress.png')}
          className="md:w-1/3 w-3/5 md:mx-0 mx-auto"
        />
        <div className="flex-1 md:pt-12 pt-6">
          <h1 className="text-2xl text-center text-black font-medium leading-relaxed tracking-wide">
            {'Never Loose Your Progress'}
          </h1>
          <h2 className="md:w-3/4 w-full mt-3 mx-auto text-lg text-center text-gray-800 font-medium leading-relaxed tracking-wide">
            {
              'Your Progress Gets Automatically Saved, So Just Pick Up Your Favourite Episode From Where You Left.'
            }
          </h2>
        </div>
      </div>

      <h1 className="md:mt-8 mt-24 text-center font-medium tracking-wide">
        <span className="text-xl">{'Not Sure Where To Start ? '}</span>
        <span className="text-lg">
          {'Check Out These Critically Acclaimed Podcasts'}
        </span>
      </h1>
      <div className="flex mt-6">
        <Grid cols={{ SM: 3, MD: 6, LG: 6 }}>
          {acclaimedPodcasts.map((p) => (
            <div key={p.id} className="flex-none md:px-5 px-3 mb-4">
              <PodcastLink podcastUrlParam={p.urlParam}>
                <a>
                  <img
                    className="w-full h-auto mb-2 flex-none object-contain rounded-lg border"
                    src={getImageUrl(p.urlParam)}
                  />
                </a>
              </PodcastLink>
              <PodcastLink podcastUrlParam={p.urlParam}>
                <a className="text-xs text-gray-900 tracking-wide font-medium leading-snug line-clamp-2">
                  {p.title}
                </a>
              </PodcastLink>
            </div>
          ))}
        </Grid>
      </div>
    </div>
  )
}

export default HomeView
