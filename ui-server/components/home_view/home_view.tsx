import classnames from 'classnames'
import Grid from 'components/grid'
import { PodcastLink } from 'components/link'
import SignInButton from 'components/sign_in_button'
import React from 'react'
import { getAssetUrl, getImageUrl } from 'utils/dom'

const acclaimedPodcasts = [
  {
    id: 'dJ8qPa',
    urlParam: 'serial-dJ8qPa',
    title: 'Serial',
    author: 'This American Life',
  },
  {
    id: 'epZEXb',
    urlParam: 'reply-all-epZEXb',
    title: 'Reply All',
    author: 'Gimlet',
  },
  {
    id: 'qaQAMa',
    urlParam: '99-invisible-qaQAMa',
    title: '99% Invisible',
    author: 'Roman Mars',
  },
  {
    id: 'en56Dd',
    urlParam: 'beautiful-stories-from-anonymous-people-en56Dd',
    title: 'Beautiful Stories From Anonymous People',
    author: 'Earwolf & Chris Gethard',
  },
  {
    id: 'dwpkwe',
    urlParam: 'the-dollop-with-dave-anthony-and-gareth-reynolds-dwpkwe',
    title: 'The Dollop with Dave Anthony and Gareth Reynolds',
    author: 'All Things Comedy | Wondery',
  },
  {
    id: 'dLZQpa',
    urlParam: 's-town-dLZQpa',
    title: 'S-Town',
    author: 'Serial & This American Life',
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
              'Discover, subscribe, curate your favourite podcasts and episodes. Sign in now to get started.'
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
              'Search for person, topic, place, podcast, episode.. From directory of 638,417 podcasts and 21,07,559 episodes.'
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
              'Subscribe to your favourite podcasts and get lastest episodes in your subscriptions feed.'
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
              'Your progress gets automatically saved, So just pick up your favourite episode from where you left off.'
            }
          </h2>
        </div>
      </div>

      <h1 className="md:mt-8 mt-24 text-center font-medium tracking-wide">
        <span className="text-xl">{'Not Sure Where To Start ? '}</span>
        <span className="text-lg">
          {'Check out these critically acclaimed podcasts'}
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
