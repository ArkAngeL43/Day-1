sudo apt-get install postgresql
sudo apt-get install libpq-dev
sudo service postgresql start
sudo update-rc.d -f postgresql remove
echo '[ + ] Creating username and owner user1'
sudo -u postgres createuser user1
sudo -u postgres createuser reaper
psql -d testdb -f log/sql/test.sql
echo '[ + ] Creating database setting owner user1'
sudo -u postgres createdb testdb --owner user1
clear 
echo '[ + ] Creating TABLE Test.sql'
psql -d testdb -f log/sql/test.sql
echo '[ + ] Editing user and owner user1'
echo "[ * ] ALTER USER user1 WITH PASSWORD 'fancy-bear-2021-fuckg';"
echo '[ * ] please enter that in the following input'
sudo -u postgres psql postgres
sudo apt install r-base
