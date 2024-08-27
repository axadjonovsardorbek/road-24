import Search from 'antd/es/input/Search'
import './style.scss'
import { Button, Table } from 'antd'
import { PlusCircleOutlined } from '@ant-design/icons'
export const Lisences = () => {

    const dataSource = [
        {
            key: '1',
            first_name: 'John',
            last_name: 'Doe',
            middle_name: 'Michael',
            date_of_birth: '1990-01-15',
            address: '123 Main St, Springfield, IL',
            date_of_issue: '2020-05-10',
            expiry_date: '2030-05-10',
            category: 'B',
            place_of_issue: 'Springfield DMV',
        },
        {
            key: '2',
            first_name: 'Jane',
            last_name: 'Smith',
            middle_name: 'Elizabeth',
            date_of_birth: '1985-03-22',
            address: '456 Oak St, Lincoln, NE',
            date_of_issue: '2019-08-15',
            expiry_date: '2029-08-15',
            category: 'C',
            place_of_issue: 'Lincoln DMV',
        },
        {
            key: '3',
            first_name: 'Robert',
            last_name: 'Johnson',
            middle_name: 'Andrew',
            date_of_birth: '1978-07-10',
            address: '789 Pine St, Madison, WI',
            date_of_issue: '2021-11-05',
            expiry_date: '2031-11-05',
            category: 'A',
            place_of_issue: 'Madison DMV',
        },
        {
            key: '4',
            first_name: 'Emily',
            last_name: 'Davis',
            middle_name: 'Marie',
            date_of_birth: '1992-12-30',
            address: '321 Elm St, Columbus, OH',
            date_of_issue: '2022-01-20',
            expiry_date: '2032-01-20',
            category: 'D',
            place_of_issue: 'Columbus DMV',
        },
        {
            key: '5',
            first_name: 'Michael',
            last_name: 'Wilson',
            middle_name: 'James',
            date_of_birth: '1988-05-18',
            address: '654 Cedar St, Denver, CO',
            date_of_issue: '2018-09-25',
            expiry_date: '2028-09-25',
            category: 'B',
            place_of_issue: 'Denver DMV',
        },
        {
            key: '6',
            first_name: 'Sarah',
            last_name: 'Brown',
            middle_name: 'Lynn',
            date_of_birth: '1995-04-12',
            address: '987 Birch St, Portland, OR',
            date_of_issue: '2021-07-14',
            expiry_date: '2031-07-14',
            category: 'B',
            place_of_issue: 'Portland DMV',
        },
        {
            key: '7',
            first_name: 'David',
            last_name: 'Martinez',
            middle_name: 'Luis',
            date_of_birth: '1983-09-23',
            address: '111 Maple St, Austin, TX',
            date_of_issue: '2017-02-03',
            expiry_date: '2027-02-03',
            category: 'C',
            place_of_issue: 'Austin DMV',
        },
        {
            key: '8',
            first_name: 'Laura',
            last_name: 'Garcia',
            middle_name: 'Maria',
            date_of_birth: '1997-11-29',
            address: '222 Cherry St, San Diego, CA',
            date_of_issue: '2022-03-18',
            expiry_date: '2032-03-18',
            category: 'D',
            place_of_issue: 'San Diego DMV',
        },
        {
            key: '9',
            first_name: 'James',
            last_name: 'Lee',
            middle_name: 'William',
            date_of_birth: '1980-06-15',
            address: '333 Oak St, Seattle, WA',
            date_of_issue: '2016-09-21',
            expiry_date: '2026-09-21',
            category: 'A',
            place_of_issue: 'Seattle DMV',
        },
        {
            key: '10',
            first_name: 'Patricia',
            last_name: 'Anderson',
            middle_name: 'Grace',
            date_of_birth: '1993-02-08',
            address: '444 Pine St, Chicago, IL',
            date_of_issue: '2020-12-10',
            expiry_date: '2030-12-10',
            category: 'B',
            place_of_issue: 'Chicago DMV',
        },
    ];



    const columns = [
        {
            title: 'First Name',
            dataIndex: 'first_name',
            key: 'first_name',
        },
        {
            title: 'Last Name',
            dataIndex: 'last_name',
            key: 'last_name',
        },
        {
            title: 'Middle Name',
            dataIndex: 'middle_name',
            key: 'middle_name',
        },
        {
            title: 'Date of Birth',
            dataIndex: 'date_of_birth',
            key: 'date_of_birth',
        },
        {
            title: "Address",
            dataIndex: "address",
            key: "address"
        },
        {
            title: "Date of Issue",
            dataIndex: "date_of_issue",
            key: "date_of_issue"
        },
        {
            title: "Expiry Date",
            dataIndex: "expiry_date",
            key: "expiry_date"
        },
        {
            title: "Category",
            dataIndex: "category",
            key: "category"
        },
        {
            title: "Place of Issue",
            dataIndex: "place_of_issue",
            key: "place_of_issue"
        },
    ]

    return (
        <div className='lisences'>
            <div className='users-header'>
                <Search className='users-search' placeholder="Search Cars..." enterButton />
                <Button type="primary" icon={<PlusCircleOutlined />}>
                    Add Lisence
                </Button>
            </div>
            <div>
                <Table dataSource={dataSource} columns={columns} />
            </div>
        </div>
    )
}
