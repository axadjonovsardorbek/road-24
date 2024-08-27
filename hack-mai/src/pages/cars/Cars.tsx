import Search from 'antd/es/input/Search'
import './style.scss'
import { Button, Table } from 'antd'
import { PlusCircleOutlined } from '@ant-design/icons'
export const Cars = () => {

    const dataSource = [
        {
            key: '1',
            car_type: 'Sedan',
            car_model: 'Toyota Camry',
            car_color: 'Black',
            car_year: '2020',
            car_vin: '1HGCM82633A123456',
            engine_number: 'EN123456789',
            horse_power: '203',
            plate_number: 'ABC-1234',
        },
        {
            key: '2',
            car_type: 'SUV',
            car_model: 'Honda CR-V',
            car_color: 'White',
            car_year: '2019',
            car_vin: '2T1BR32E29C123456',
            engine_number: 'EN987654321',
            horse_power: '190',
            plate_number: 'XYZ-5678',
        },
        {
            key: '3',
            car_type: 'Coupe',
            car_model: 'Ford Mustang',
            car_color: 'Red',
            car_year: '2021',
            car_vin: '1ZVBP8AM8D5234567',
            engine_number: 'EN112233445',
            horse_power: '450',
            plate_number: 'MUS-1234',
        },
        {
            key: '4',
            car_type: 'Truck',
            car_model: 'Chevrolet Silverado',
            car_color: 'Blue',
            car_year: '2018',
            car_vin: '3GCPKTE73BG123456',
            engine_number: 'EN998877665',
            horse_power: '355',
            plate_number: 'TRK-5678',
        },
        {
            key: '5',
            car_type: 'Hatchback',
            car_model: 'Volkswagen Golf',
            car_color: 'Silver',
            car_year: '2022',
            car_vin: 'WVWZZZ1KZAW123456',
            engine_number: 'EN556677889',
            horse_power: '147',
            plate_number: 'VWG-1234',
        },
        {
            key: '6',
            car_type: 'Convertible',
            car_model: 'BMW Z4',
            car_color: 'Yellow',
            car_year: '2020',
            car_vin: 'WBALM5C52BE123456',
            engine_number: 'EN334455667',
            horse_power: '382',
            plate_number: 'BMW-5678',
        },
        {
            key: '7',
            car_type: 'Minivan',
            car_model: 'Honda Odyssey',
            car_color: 'Gray',
            car_year: '2017',
            car_vin: '5FNRL5H68CB123456',
            engine_number: 'EN445566778',
            horse_power: '280',
            plate_number: 'VAN-9876',
        },
        {
            key: '8',
            car_type: 'Pickup',
            car_model: 'Ram 1500',
            car_color: 'Green',
            car_year: '2019',
            car_vin: '1C6RR7KT3HS123456',
            engine_number: 'EN667788990',
            horse_power: '395',
            plate_number: 'RAM-5432',
        },
        {
            key: '9',
            car_type: 'Crossover',
            car_model: 'Nissan Rogue',
            car_color: 'Brown',
            car_year: '2021',
            car_vin: 'JN8AT2MV7LW123456',
            engine_number: 'EN778899001',
            horse_power: '181',
            plate_number: 'NSN-6543',
        },
        {
            key: '10',
            car_type: 'Electric',
            car_model: 'Tesla Model 3',
            car_color: 'Blue',
            car_year: '2022',
            car_vin: '5YJ3E1EA7LF123456',
            engine_number: 'EN123450987',
            horse_power: '283',
            plate_number: 'TES-1234',
        },
    ];


    const columns = [
        {
            title: 'Type',
            dataIndex: 'car_type',
            key: 'car_type',
        },
        {
            title: 'Model',
            dataIndex: 'car_model',
            key: 'car_model',
        },
        {
            title: 'Color',
            dataIndex: 'car_color',
            key: 'car_color',
        },
        {
            title: "Year",
            dataIndex: "car_year",
            key: "car_year"
        },
        {
            title: "VIN",
            dataIndex: "car_vin",
            key: "car_vin"
        },
        {
            title: "Engine Number",
            dataIndex: "engine_number",
            key: "engine_number"
        },
        {
            title: "Horse Power",
            dataIndex: "horse_power",
            key: "horse_power"
        },
        {
            title: "Plate ",
            dataIndex: "plate_number",
            key: "plate_number"
        },
    ]

    return (
        <div className='users'>
            <div className='users-header'>
                <Search className='users-search' placeholder="Search Cars..." enterButton />
                <Button type="primary" icon={<PlusCircleOutlined />}>
                    Add Car
                </Button>
            </div>
            <div>
                <Table dataSource={dataSource} columns={columns} />
            </div>
        </div>
    )
}
