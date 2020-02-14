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

const HomeView: React.FC<{}> = () => {
  return (
    <div>
      <div className="flex md:flex-row flex-col pt-6">
        <div className="flex-1 pt-16">
          <h1 className="text-5xl text-center text-purple-700 font-semibold tracking-wide">
            {'Phenopod'}
          </h1>
          <h2 className="text-lg text-center text-gray-900 font-medium leading-relaxed tracking-wide">
            {'Simple, Yet Powerful Podcast Player'}
          </h2>
          <h3 className="w-3/4 mt-6 mx-auto text-sm text-center text-gray-800 font-medium tracking-wide">
            {
              'Discover, Subscribe, Curate Your Favourite Podcasts And Episodes. Sign In Now To Get Started.'
            }
          </h3>
          <div className="w-56 h-8 mx-auto mt-3">
            <SignInButton />
          </div>
        </div>
        <img src={getAssetUrl('listen-to-podcasts.svg')} className="w-1/2" />
      </div>

      <div className="flex mt-12 pl-6">
        <img src={getAssetUrl('powerful-search.png')} className="w-1/3" />
        <div className="flex-1 pt-14">
          <h1 className="text-2xl text-center text-black font-medium leading-relaxed tracking-wide">
            {'Powerful Search Engine'}
          </h1>
          <h2 className="w-3/4 mt-3 mx-auto text-lg text-center text-gray-800 font-medium leading-relaxed tracking-wide">
            {
              'Search For Person, Topic, Place, Podcast, Episode.. From A Directory Of 638,417 Podcasts And 21,07,559 Episodes.'
            }
          </h2>
        </div>
      </div>

      <div className="flex mt-16 pl-6">
        <img src={getAssetUrl('subscriptions-feed.png')} className="w-1/3" />
        <div className="flex-1 pt-12">
          <h1 className="text-2xl text-center text-black font-medium leading-relaxed tracking-wide">
            {'Subscriptions'}
          </h1>
          <h2 className="w-3/4 mt-3 mx-auto text-lg text-center text-gray-800 font-medium leading-relaxed tracking-wide">
            {
              'Subscribe To Your Favourite Podcasts And Get Lastest Episodes In Your Subscriptions Feed.'
            }
          </h2>
        </div>
      </div>

      <div className="flex mt-16 pl-6">
        <img src={getAssetUrl('sync-progress.png')} className="w-1/3" />
        <div className="flex-1 pt-14">
          <h1 className="text-2xl text-center text-black font-medium leading-relaxed tracking-wide">
            {'Never Loose Your Progress'}
          </h1>
          <h2 className="w-3/4 mt-3 mx-auto text-lg text-center text-gray-800 font-medium leading-relaxed tracking-wide">
            {
              'Your Progress Gets Automatically Saved, So Just Pick Up Your Favourite Episode From Where You Left.'
            }
          </h2>
        </div>
      </div>

      <h1 className="mt-8 text-center font-medium tracking-wide">
        <span className="text-xl">
          {'Not Able to Decide Where To Start ? '}
        </span>
        <span className="text-lg">
          {'Check Out These Critically Acclaimed Podcasts'}
        </span>
      </h1>
      <div className="flex mt-6">
        {acclaimedPodcasts.map((p) => (
          <div key={p.id} className="flex-none w-1/6 px-5 mb-4">
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
      </div>
    </div>
  )
}

export default HomeView
