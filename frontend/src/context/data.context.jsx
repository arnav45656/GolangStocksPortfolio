import { createContext, useEffect, useState } from 'react';
import { fetchNews } from '../utils/fetchNews';
import { stocksArr } from '../utils/seeds/stocksArr';
import { getUser, createUser } from '../api/api';

const DataContext = createContext();

function DataProviderWrapper(props) {
    const [trendingStocks, setTrendingStocks] = useState([]);
    const [buttonSelected, setButtonSelected] = useState('all');
    const [news, setNews] = useState([]);
    const [wallet, setWallet] = useState({ deposit: 0, investment: 0 });
    const [investments, setInvestments] = useState([]);
    const [sells, setSells] = useState([]);
    const [balance, setBalance] = useState(0);
    const [user, setUser] = useState({ userName: '' });
    const [watchlist, setWatchlist] = useState([]);
    const [searchResults, setSearchResults] = useState([]);
    const [isSearching, setIsSearching] = useState(false);

    // useEffect(() => {
    //     const loadData = async () => {
    //         if (user.userName) {
    //             getTrendingStocks();
    //             setDefaultWatchlist();
    //             await loadUserData(user.userName);
    //             const newsResponse = await fetchNews();
    //             setNews(newsResponse);
                
    //         }
    //     };

    //     loadData();
    // }, [user.userName]);

    // const loadUserData = async (username) => {
    //     if (!username) {
    //         console.error('Username is undefined, skipping user data load.');
    //         return;
    //     }

    //     try {
    //         const userData = await getUser(username);
    //         console.log('User data loaded:', userData);
    //         setWallet(userData.Wallet || { deposit: 0, investment: 0 });
    //         setInvestments(userData.Investments || []);
    //         setSells(userData.Sells || []);
    //         setBalance(userData.Balance || 0);
    //         setWatchlist(userData.Watchlist || []);
            
    //         setUser({ userName: userData.UserName});
    //     } catch (error) {
    //         if (error.response && error.response.status === 404) {
    //             const newUser = {
    //                 userName: username,
    //                 wallet,
    //                 investments,
    //                 sells,
    //                 balance,
    //                 watchlist,
    //             };
    //             try {
    //                 const createdUser = await createUser(newUser);
    //                 setUser(createdUser);
    //             } catch (createError) {
    //                 console.error('Error creating user:', createError);
    //             }
    //         } else {
    //             console.error('Error loading user data:', error);
    //         }
    //     }
    // };

    // 
    

    useEffect(() => {
        if (user.userName && user.userName !== '') {
            console.log(user)
          loadUserData(user.userName);
          fetchNews().then(response => {
            setNews(response);
          });
          getTrendingStocks();
          setDefaultWatchlist();
        }
      }, [user.userName]);
    
      const loadUserData = (username) => {
        getUser(username).then(userData => {
          
          setInvestments(userData.Investments || []);
          setSells(userData.Sells || []);
          setBalance(userData.Balance || 0);
          
          
        }).catch(error => {
            console.log(error)
          if (username && error.code === "ERR_BAD_REQUEST") {
            console.log(username)
            const newUser = {
              userName: username,
              wallet,
              investments :[],
              sells : [],
              balance : balance,
              watchlist: [], 
            };
            createUser(newUser).then(createdUser => {
                console.log(createdUser)
              setUser({ userName: createdUser.UserName });
              
            }).catch(createError => {
              console.error('Error creating user:', createError);
            });
          } else {
            console.error('Error loading user data:', error);
          }
        });
      };

    useEffect(() => {
        if (investments.length > 0) {
            const lastStock = investments[investments.length - 1]
            if (lastStock.Type === 'buy') {
                setBalance(balance - Number(lastStock.Amount))
            }
            if (lastStock.Type === 'sell') {
                setBalance(balance + Number(lastStock.Amount))
            }
        }
    }, [investments])

    useEffect(() => {
        if (sells) {
            const reducedSells = [...sells].reduce(function (acc, investment) { return acc + Number(investment.amount) }, 0)
            const currentBalance = Number(balance) + Number(reducedSells)
            setBalance(currentBalance.toFixed(2))
        }
    }, [sells])


    const getTrendingStocks = (type) => {
        const winners = stocksArr.filter(stock => !stock.PriceChange.includes("-"));
        const lossers = stocksArr.filter(stock => stock.PriceChange.includes("-") && stock.PriceChange !== "-");
        const sortedWinners = winners.sort((a, b) => b.PriceChange.localeCompare(a.PriceChange));
        const sortedLossers = lossers.sort((a, b) => b.PriceChange.localeCompare(a.PriceChange));
        const popular = [...sortedWinners.slice(0, 5), ...sortedLossers.slice(0, 5)].sort((a, b) => a.CompanyName.localeCompare(b.CompanyName));

        if (type === 'winners') {
            setTrendingStocks(sortedWinners.slice(0, 5));
            setButtonSelected('winners');
        } else if (type === 'lossers') {
            setTrendingStocks(sortedLossers.slice(0, 5));
            setButtonSelected('lossers');
        } else {
            setTrendingStocks(popular);
            setButtonSelected('all');
        }
    };

    const setDefaultWatchlist = () => {
        const selectedStocks = stocksArr.filter(stock =>
            stock.Symbol === 'AAPL' || stock.Symbol === 'TSLA' || stock.Symbol === 'MSFT'
        );
        setWatchlist(selectedStocks);
    };

    return (
        <DataContext.Provider
            value={{
                trendingStocks,
                setTrendingStocks,
                watchlist,
                setWatchlist,
                getTrendingStocks,
                buttonSelected,
                news,
                setNews,
                wallet,
                setWallet,
                investments,
                setInvestments,
                balance,
                setBalance,
                user,
                setUser,
                sells,
                setSells,
                searchResults,
                setSearchResults,
                isSearching,
                setIsSearching,
            }}
        >
            {props.children}
        </DataContext.Provider>
    );
}

export { DataContext, DataProviderWrapper };
